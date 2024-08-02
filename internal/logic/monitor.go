package logic

import (
	"context"
	"github.com/lucsky/cuid"
	"heimdall/internal/dal"
	"heimdall/internal/dal/model"
	"heimdall/internal/dep"
	"heimdall/internal/service/github"
	"heimdall/pkg/function"
	"log"
	"strings"
	"sync"
	"time"
)

type Monitor struct {
	RepoDAL   dal.IGitRepositoryDAL
	CommitDAL dal.ICommitDAL
	GitHub    github.IRepositoryService
}

func NewMonitorLogic(dependencies *dep.Dependencies) *Monitor {
	return &Monitor{
		RepoDAL:   dependencies.DAL.GitRepositoryDAl,
		CommitDAL: dependencies.DAL.CommitDAL,
		GitHub:    dependencies.GitHub,
	}
}

var rcMu sync.Mutex

func (m *Monitor) RetrieveCommit() {

	rcMu.Lock()
	defer rcMu.Unlock()

	var (
		jobs   = make(chan model.GitRepository, 5)
		wg     sync.WaitGroup
		cursor = ""
	)

	// setup worker pool
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go m.processCommitAddition(&wg, jobs)
	}

	for {

		repos, err := m.RepoDAL.ListRepoCursor(context.Background(), cursor, 10)
		if err != nil {

		}
		if len(repos) == 0 { // exit loop if no more repositories exist
			break
		}

		for _, repo := range repos {
			jobs <- repo
		}

		cursor = repos[len(repos)-1].Id

	}
	close(jobs)

	wg.Wait()
}

func (m *Monitor) processCommitAddition(wg *sync.WaitGroup, repoJob chan model.GitRepository) {

	defer wg.Done()
	for repo := range repoJob {

		var pullFrom time.Time
		var err error
		if err := function.Retry(3, time.Second*2, func() error {
			pullFrom, err = m.CommitDAL.RepoLastCommitDate(context.Background(), repo.Id)
			return err
		}); err != nil {
			log.Printf("Failed to get last commit date for repo %s: %s", repo.Id, err.Error())
			continue
		}

		var (
			commits []github.Commit
			link    = ""
		)

		for {

			commits, link, err = m.GitHub.ListCommit(repo.Name, repo.Name, pullFrom, link)
			if err != nil {
				break
			}
			if len(commits) == 0 {
				break
			}

			newCommits := make([]model.Commit, len(commits))
			for _, commit := range commits {
				newCommits = append(newCommits, model.Commit{
					Id:      cuid.New(),
					RepoId:  repo.Id,
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

			if err := function.Retry(3, time.Second*2, func() error {
				return m.CommitDAL.AddCommits(context.Background(), newCommits)
			}); err != nil {
				log.Printf("[Commit.Addition] failed to add commits from %s for %s\n", link, repo.Id)
			}

			if link == "" {
				break
			}

			nextLink := strings.Split(strings.Split(link, ",")[0], ";")[0]
			nextLink = strings.TrimPrefix(nextLink, "<")
			link = strings.TrimSuffix(nextLink, ">")

		}

	}

}
