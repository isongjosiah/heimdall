package tests

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lucsky/cuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"heimdall/internal/api/rest"
	"heimdall/internal/config"
	"heimdall/internal/dal"
	"heimdall/internal/dal/dalfakes"
	"heimdall/internal/dal/model"
	"heimdall/internal/dep"
	"heimdall/internal/logic"
	"heimdall/internal/service/github"
	"heimdall/internal/service/github/githubfakes"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	restApi *rest.API

	commitDal     *dalfakes.FakeICommitDAL
	repositoryDal *dalfakes.FakeIGitRepositoryDAL

	githubService *githubfakes.FakeIRepositoryService

	dependencies  *dep.Dependencies
	serverHandler http.Handler

	dbMock sqlmock.Sqlmock
)

func TestMain(m *testing.M) {

	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Print("error setting up mock database", err)
		os.Exit(1)
	}

	dbMock = mock
	commitDal = &dalfakes.FakeICommitDAL{}
	repositoryDal = &dalfakes.FakeIGitRepositoryDAL{}
	githubService = &githubfakes.FakeIRepositoryService{}

	dependencies = &dep.Dependencies{
		DAL: &dal.DAL{
			SqlDB:            bun.NewDB(db, pgdialect.New()),
			CommitDAL:        commitDal,
			GitRepositoryDAl: repositoryDal,
		},
		GitHub: githubService,
	}

	restApi = &rest.API{
		Config: &config.Config{
			GithubToken: "fakeToken",
		},
		Dep:   dependencies,
		Logic: logic.New(dependencies),
	}

	serverHandler = restApi.SetupServerHandler()
	os.Exit(m.Run())
}

func TestMonitor_RetrieveCommit(t *testing.T) {

	databaseRepoRes := []model.GitRepository{
		{
			Id:             cuid.New(),
			Name:           "fakeName",
			Owner:          "fakeOwner",
			Description:    "fakeDescription",
			URL:            "fakeUrl",
			Language:       "fakeLanguageInformation",
			ForkCount:      400,
			StarsCount:     790,
			OpenIssueCount: 800,
			WatchersCount:  560,
			CreatedAt:      time.Now(),
			AddedAt:        time.Now(),
			UpdatedAt:      time.Now(),
		},
		{
			Id:             cuid.New(),
			Name:           "fakeName2",
			Owner:          "fakeOwner2",
			Description:    "fakeDescription2",
			URL:            "fakeUrl2",
			Language:       "fakeLanguageInformation2",
			ForkCount:      400,
			StarsCount:     790,
			OpenIssueCount: 800,
			WatchersCount:  560,
			CreatedAt:      time.Now(),
			AddedAt:        time.Now(),
			UpdatedAt:      time.Now(),
		},
	}
	listCommitResponse := []github.Commit{
		{
			URL:         "https://api.github.com/repos/user/repo/commits/1",
			SHA:         "sha1",
			CommentsURL: "https://api.github.com/repos/user/repo/commits/1/comments",
			Commit: struct {
				URL    string `json:"url"`
				Author struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"author"`
				Committer struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"committer"`
				Message      string `json:"message"`
				CommentCount int    `json:"comment_count"`
			}{
				URL: "https://api.github.com/repos/user/repo/git/commits/1",
				Author: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Author1",
					Email: "author1@example.com",
					Date:  time.Now().AddDate(0, 0, -1),
				},
				Committer: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Committer1",
					Email: "committer1@example.com",
					Date:  time.Now().AddDate(0, 0, -1),
				},
				Message:      "Commit message 1",
				CommentCount: 10,
			},
		},
		{
			URL:         "https://api.github.com/repos/user/repo/commits/2",
			SHA:         "sha2",
			CommentsURL: "https://api.github.com/repos/user/repo/commits/2/comments",
			Commit: struct {
				URL    string `json:"url"`
				Author struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"author"`
				Committer struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"committer"`
				Message      string `json:"message"`
				CommentCount int    `json:"comment_count"`
			}{
				URL: "https://api.github.com/repos/user/repo/git/commits/2",
				Author: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Author2",
					Email: "author2@example.com",
					Date:  time.Now().AddDate(0, 0, -2),
				},
				Committer: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Committer2",
					Email: "committer2@example.com",
					Date:  time.Now().AddDate(0, 0, -2),
				},
				Message:      "Commit message 2",
				CommentCount: 8,
			},
		},
		{
			URL:         "https://api.github.com/repos/user/repo/commits/3",
			SHA:         "sha3",
			CommentsURL: "https://api.github.com/repos/user/repo/commits/3/comments",
			Commit: struct {
				URL    string `json:"url"`
				Author struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"author"`
				Committer struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"committer"`
				Message      string `json:"message"`
				CommentCount int    `json:"comment_count"`
			}{
				URL: "https://api.github.com/repos/user/repo/git/commits/3",
				Author: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Author3",
					Email: "author3@example.com",
					Date:  time.Now().AddDate(0, 0, -3),
				},
				Committer: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Committer3",
					Email: "committer3@example.com",
					Date:  time.Now().AddDate(0, 0, -3),
				},
				Message:      "Commit message 3",
				CommentCount: 5,
			},
		},
		{
			URL:         "https://api.github.com/repos/user/repo/commits/4",
			SHA:         "sha4",
			CommentsURL: "https://api.github.com/repos/user/repo/commits/4/comments",
			Commit: struct {
				URL    string `json:"url"`
				Author struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"author"`
				Committer struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				} `json:"committer"`
				Message      string `json:"message"`
				CommentCount int    `json:"comment_count"`
			}{
				URL: "https://api.github.com/repos/user/repo/git/commits/4",
				Author: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Author4",
					Email: "author4@example.com",
					Date:  time.Now().AddDate(0, 0, -4),
				},
				Committer: struct {
					Name  string    `json:"name"`
					Email string    `json:"email"`
					Date  time.Time `json:"date"`
				}{
					Name:  "Committer4",
					Email: "committer4@example.com",
					Date:  time.Now().AddDate(0, 0, -4),
				},
				Message:      "Commit message 4",
				CommentCount: 12,
			},
		},
	}

	repositoryDal.ListRepoCursorReturns(databaseRepoRes, nil)
	repositoryDal.ListRepoCursorReturnsOnCall(1, nil, nil)
	githubService.ListCommitReturns(listCommitResponse, "", nil)
	commitDal.AddCommitsReturns(nil)

	restApi.Logic.Monitor.RetrieveCommit()
	assert.Equal(t, 2, repositoryDal.ListRepoCursorCallCount())
	assert.Equal(t, 2, githubService.ListCommitCallCount())
	assert.Equal(t, 2, commitDal.AddCommitsCallCount())

}
