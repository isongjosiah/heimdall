package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/uptrace/bun"
	"heimdall/internal/dal"
	"heimdall/internal/dal/model"
	"heimdall/internal/dep"
	"heimdall/internal/service/github"
	"heimdall/internal/service/queue"
	"heimdall/internal/value"
	"heimdall/pkg/function"
	"log"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Repository struct {
	SqlDB     *bun.DB
	GitHub    github.IRepositoryService
	RepoDAL   dal.IGitRepositoryDAL
	CommitDAL dal.ICommitDAL
}

func NewRepositoryLogic(dep *dep.Dependencies) *Repository {
	return &Repository{
		SqlDB:     dep.DAL.SqlDB,
		RepoDAL:   dep.DAL.GitRepositoryDAl,
		CommitDAL: dep.DAL.CommitDAL,
		GitHub:    dep.GitHub,
	}
}

// Commits retrieves a paginated list of commit for a repository
func (r *Repository) Commits(ctx context.Context, repoName string, queryValue url.Values) (model.CommitList, string, string, error) {

	repository, message, err := r.RepoDAL.RepoByName(ctx, repoName)
	if err != nil {
		return model.CommitList{}, "", message, err
	}

	commitList, err := r.CommitDAL.CommitsByRepoId(ctx, repository.Id, queryValue)
	if err != nil {
		return model.CommitList{}, "", message, err
	}

	return commitList, value.Success, "Repository commit retrieved", nil

}

// AddNewRepository ...
func (r *Repository) AddNewRepository(ctx context.Context, payload model.AddRepositoryInput) (string, string, error) {

	err := (queue.RMQProducer{
		Queue: "pull-commit",
	}).PublishMessage(payload)
	if err != nil {
		return value.Error, "Unable to process request. Please try again later", err
	}

	return value.Success, "Repository and commits are being added asynchronously", nil

}

func (r *Repository) handleRepositoryAddition(ctx context.Context, delivery amqp091.Delivery) {

	// parse data
	var repo model.AddRepositoryInput
	err := json.Unmarshal(delivery.Body, &repo)
	if err != nil {
		queue.Nack(delivery)
	}

	message := "Repository added successfully"
	err = r.SqlDB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		exists, err := r.RepoDAL.RepoExists(ctx, repo.RepoName)
		if err != nil {
			message = "Unable to check if repository exists" + err.Error()
			return err
		}
		if exists {
			message = "Repository already exists"
			return nil
		}

		gitRepo, err := r.GitHub.GetRepository(repo.RepoOwner, repo.RepoName)
		if err != nil {
			message = "Unable to retrieve git repository: " + err.Error()
			return err
		}

		newRepo := model.GitRepository{
			Id:              cuid.New(),
			Name:            gitRepo.Name,
			Owner:           repo.RepoOwner,
			Description:     repo.RepoName,
			URL:             gitRepo.URL,
			Language:        "",
			ForkCount:       gitRepo.ForksCount,
			StarsCount:      gitRepo.StargazersCount,
			OpenIssueCount:  gitRepo.OpenIssuesCount,
			WatchersCount:   gitRepo.WatchersCount,
			InitialPullDone: true,
			CreatedAt:       time.Now(),
			AddedAt:         time.Now(),
			UpdatedAt:       time.Now(),
		}
		if err := r.RepoDAL.AddRepo(ctx, newRepo); err != nil {
			message = "failed to persist repository metadata" + err.Error()
			return err
		}
		repo.RepoId = newRepo.Id

		if err := (queue.RMQProducer{
			Queue: "pull-commit",
		}).PublishMessage(repo); err != nil {
			message = "failed to push repo detail to pull-commit queue"
			return err
		}

		return nil
	})
	if err != nil {
		deliveryCount, _ := delivery.Headers["x-delivery-count"].(int32)
		if deliveryCount > 3 {
			log.Println(message)
			queue.Ack(delivery)
		}
		queue.Nack(delivery)
	}

	queue.Ack(delivery)

}

type InitialPullJob struct {
	repoId     string
	commits    []github.Commit
	commitLink string
}

func (r *Repository) handleInitialPull(ctx context.Context, delivery amqp091.Delivery) {

	// parse data
	var repo model.AddRepositoryInput
	err := json.Unmarshal(delivery.Body, &repo)
	if err != nil {
		queue.Nack(delivery)
	}

	var (
		wg          sync.WaitGroup
		commits     []github.Commit
		jobs        = make(chan InitialPullJob, 5)
		link        = ""
		insertCount atomic.Int32
	)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go r.processCommitAddition(&wg, jobs, &insertCount)
	}

	for {

		job := InitialPullJob{
			commitLink: link,
		}

		commits, link, err = r.GitHub.ListCommit(repo.RepoOwner, repo.RepoName, repo.FetchCommitFrom, link)
		if err != nil {
			queue.Ack(delivery)
			break
		}
		if len(commits) == 0 {
			break
		}

		job.repoId = repo.RepoId
		job.commits = commits
		jobs <- job

		if link == "" {
			break
		}

		nextLink := strings.Split(strings.Split(link, ",")[1], ";")[0]
		nextLink = strings.TrimPrefix(nextLink, "<")
		link = strings.TrimSuffix(nextLink, ">")

	}

	close(jobs)
	wg.Wait()
	if insertCount.Load() > 0 {
		if err := function.Retry(10, time.Second*2, func() error {
			return r.RepoDAL.UpdateRepo(context.Background(), repo.RepoId, map[string]any{"initial_pull_done": true})
		}); err != nil {
			log.Println("[Commit.Addition] failed to update repository status for initial commit " + err.Error())
		}
	}
}

func (r *Repository) processCommitAddition(wg *sync.WaitGroup, commitJob chan InitialPullJob, insertCount *atomic.Int32) {
	defer wg.Done()
	for commits := range commitJob {

		newCommits := make([]model.Commit, len(commits.commits))
		for _, commit := range commits.commits {
			newCommits = append(newCommits, model.Commit{
				Id:      cuid.New(),
				RepoId:  commits.repoId,
				Sha:     commit.SHA,
				Message: commit.Commit.Message,
				Author: model.CommitAuthor{
					Name:  commit.Commit.Author.Name,
					Email: commit.Commit.Author.Email,
				},
				Date:    commit.Commit.Author.Date,
				URL:     commit.URL,
				AddedAt: time.Now(),
			})
		}
		if err := function.Retry(4, time.Second*2, func() error {
			return r.CommitDAL.AddCommits(context.Background(), newCommits)
		}); err != nil {
			log.Printf("[Commit.Addition] failed to add commits from %s for %s\n", commits.commitLink, commits.repoId)
			continue
		}
		insertCount.Add(1)
	}
}

func (r *Repository) ResetRepositoryCollection(ctx context.Context, payload model.ResetCollectionInput) (string, string, error) {

	// check that the repository exists
	repo, status, err := r.RepoDAL.RepoByName(ctx, payload.RepoName)
	if err != nil {
		msg := "Unable to validate repository existence"
		if errors.Is(err, sql.ErrNoRows) {
			msg = "Repository does not exist"
		}
		return status, msg, err
	}

	err = r.SqlDB.RunInTx(context.Background(), &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		// delete all commits recorded
		ctx = context.WithValue(ctx, "sql-tx", tx)
		if err := r.CommitDAL.DeleteRepoCommits(ctx, payload.RepoName); err != nil {
			return err
		}

		// update commit date
		return r.RepoDAL.UpdateRepo(ctx, repo.Id, map[string]any{"pull_from": payload.StartFrom})
	})
	if err != nil {
		return value.Error, "Failed to reset repository collection", err
	}

	return value.Success, "repository collection reset successfully", nil
}
