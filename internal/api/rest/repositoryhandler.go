package rest

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"heimdall/internal/dal/model"
	"heimdall/internal/value"
	"heimdall/pkg/function"
	"net/http"
)

func (a *API) RepositoryRoutes() chi.Router {

	repositoryRouter := chi.NewRouter()
	repositoryRouter.Method(http.MethodPost, "/", Handler(a.AddRepositoryToMonitorH))
	repositoryRouter.Method(http.MethodGet, "/commits", Handler(a.GetRepositoryCommitsH))
	repositoryRouter.Method(http.MethodGet, "/reset-collection-date", Handler(a.ResetCollectionDateH))

	return repositoryRouter

}

// AddRepositoryToMonitorH handles request to add a repository to monitor to the database
func (a *API) AddRepositoryToMonitorH(w http.ResponseWriter, r *http.Request) *ServerResponse {

	var payload model.AddRepositoryInput
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return RespondWithError(err, "Bad request body", value.BadRequest, function.StatusCode(value.BadRequest))
	}

	status, message, err := a.Logic.Repository.AddNewRepository(context.Background(), payload)
	if err != nil {
		return RespondWithError(err, message, status, function.StatusCode(status))
	}

	return &ServerResponse{
		Message:    message,
		Status:     status,
		StatusCode: http.StatusCreated,
	}
}

// GetRepositoryCommitsH handles request to retrieve paginated commit for an existing repository
func (a *API) GetRepositoryCommitsH(w http.ResponseWriter, r *http.Request) *ServerResponse {

	// retrieve filter information
	queryValues := r.URL.Query()

	repoName := queryValues.Get("repo-name")
	if repoName == "" {
		return RespondWithError(errors.New("invalid parameter query value"), "repo-name is required",
			value.BadRequest, http.StatusBadRequest)
	}

	repoCommits, status, message, err := a.Logic.Repository.Commits(context.Background(), repoName, queryValues)
	if err != nil {
		return RespondWithError(err, message, status, function.StatusCode(status))
	}

	return &ServerResponse{
		Message:    message,
		Status:     status,
		StatusCode: function.StatusCode(status),
		Payload:    repoCommits,
	}
}

// ResetCollectionDateH ...
func (a *API) ResetCollectionDateH(w http.ResponseWriter, r *http.Request) *ServerResponse {

	var payload model.ResetCollectionInput
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return RespondWithError(err, "Bad request body", value.BadRequest, function.StatusCode(value.BadRequest))
	}

	status, message, err := a.Logic.Repository.ResetRepositoryCollection(context.TODO(), payload)
	if err != nil {
		return RespondWithError(err, message, status, function.StatusCode(value.BadRequest))
	}
	return &ServerResponse{
		Message:    message,
		Status:     status,
		StatusCode: http.StatusOK,
	}
}
