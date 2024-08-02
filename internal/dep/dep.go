package dep

import (
	"heimdall/internal/config"
	"heimdall/internal/dal"
	"heimdall/internal/service/github"
	"sync"
)

var (
	dep  *Dependencies
	once sync.Once
)

type Dependencies struct {
	// Services
	GitHub github.IRepositoryService

	//DAL
	DAL *dal.DAL
}

// New initializes the dependencies required for
// the application to function
func New(cfg *config.Config) *Dependencies {

	once.Do(func() {

		dep = &Dependencies{
			DAL:    dal.NewDAL(cfg),
			GitHub: github.NewService(cfg),
		}

	})

	return dep
}
