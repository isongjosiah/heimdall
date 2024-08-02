package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"heimdall/pkg/function"
	"net/http"
)

func (a *API) CommitRoutes() chi.Router {

	commitRouter := chi.NewRouter()
	commitRouter.Method(http.MethodGet, "/top-contributors", Handler(a.TopContributorsH))

	return commitRouter

}

func (a *API) TopContributorsH(w http.ResponseWriter, r *http.Request) *ServerResponse {

	queryValues := r.URL.Query()
	authorCount := function.StringToInt(queryValues.Get("author-count"))
	if authorCount == 0 { // use default of 10
		authorCount = 10
	}

	page := function.StringToInt(queryValues.Get("page"))
	if page == 0 { // use default of 10
		page = 1
	}

	topContributors, status, message, err := a.Logic.Commit.TopContributors(context.Background(), authorCount, page)
	if err != nil {
		return RespondWithError(err, message, status, function.StatusCode(status))
	}

	return &ServerResponse{
		Message:    message,
		Status:     status,
		StatusCode: http.StatusOK,
		Payload:    topContributors,
	}
}
