package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"heimdall/internal/config"
	"heimdall/internal/dep"
	"heimdall/internal/logic"
	"heimdall/internal/value"
	"log"
	"net/http"
	"strings"
	"time"
)

type API struct {
	Server *http.Server
	Config *config.Config
	Dep    *dep.Dependencies
	Logic  *logic.Logic
}

// Serve starts the http server
func (a *API) Serve() error {
	a.Server = &http.Server{
		Addr:           fmt.Sprintf(":%d", a.Config.HttpPort),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   60 * time.Second,
		Handler:        a.SetupServerHandler(),
		MaxHeaderBytes: 1024 * 1024,
	}

	log.Println("[Heimdall] Listening on", a.Config.HttpPort)
	return a.Server.ListenAndServe()
}

// Shutdown stops the http server
func (a *API) Shutdown() error {
	return a.Server.Shutdown(context.Background())
}

type Handler func(w http.ResponseWriter, r *http.Request) *ServerResponse

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h(w, r)

	if response.StatusCode == 0 && response.Err == nil {
		response.StatusCode = http.StatusOK
	}

	var responseBytes []byte
	var marshalErr error

	if response.Err != nil {
		message := response.Message
		if response.Status == value.Error {
			message = strings.Split(message, ":")[0]
		}
		responseBytes, marshalErr = json.Marshal(
			ErrorResponse{
				ErrorMessage: message,
				Status:       response.Status,
				StatusCode:   response.StatusCode,
			},
		)
	} else {
		responseBytes, marshalErr = json.Marshal(response)
	}

	if marshalErr != nil {
		// marshallErr type
	}

	WriteJSONResponse(w, response.StatusCode, responseBytes)
}

// SetupServerHandler configures the server handlers
func (a *API) SetupServerHandler() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(120 * time.Second))
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{ health: "ok" }`)); err != nil {
			logger, _ := zap.NewProduction()
			logger.Info("Unable to write health response")
		}
	})

	mux.Mount("/commits", a.CommitRoutes())
	mux.Mount("/repositories", a.RepositoryRoutes())

	return mux
}
