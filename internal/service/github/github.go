package github

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"heimdall/internal/config"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//counterfeiter:generate . IRepositoryService
type IRepositoryService interface {
	GetRepository(repoOwner, repoName string) (Repository, error)
	ListCommit(repoOwner, repoName string, startAt time.Time, link string) ([]Commit, string, error)
}

var (
	service = new(Service)
	once    sync.Once
)

type Service struct {
	authToken string
}

func NewService(config *config.Config) *Service {
	once.Do(func() {
		service = &Service{
			authToken: config.GithubToken,
		}
	})

	return service
}

// ListCommit retrieves commits from the GitHub API
func (s Service) ListCommit(repoOwner, repoName string, startAt time.Time, link string) ([]Commit, string, error) {

	path := fmt.Sprintf("/repos/%s/%s/commits", repoOwner, repoName)
	endpoint := BaseUrl + path

	if !startAt.IsZero() {
		endpoint += "?since=" + url.QueryEscape(startAt.Format(time.RFC3339))
	}
	if link != "" {
		if _, err := url.ParseRequestURI(link); err != nil {
			log.Println("[Github.ListCommit] invalid link " + link)
		} else {
			endpoint = link
		}
	}
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("Authorization", "Bearer "+s.authToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}

	if res.StatusCode != http.StatusOK {
		return nil, "", errors.New("")
	}

	var payload []Commit
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return payload, "", err
	}

	return payload, res.Header.Get("Link"), nil

}

// GetRepository ...
func (s Service) GetRepository(repoOwner, repoName string) (Repository, error) {

	path := fmt.Sprintf("/repos/%s/%s", repoOwner, repoName)
	endpoint := BaseUrl + path

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return Repository{}, err
	}
	req.Header.Set("Authorization", "Bearer "+s.authToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Repository{}, err
	}

	if res.StatusCode != http.StatusOK {
		return Repository{}, errors.New("")
	}

	var repository Repository
	if er := json.NewDecoder(res.Body).Decode(&repository); er != nil {
		return repository, err
	}

	return repository, nil
}
