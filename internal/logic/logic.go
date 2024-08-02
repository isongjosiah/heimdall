package logic

import (
	"heimdall/internal/dep"
	"sync"
)

var (
	logic = new(Logic)
	once  sync.Once
)

type Logic struct {
	Commit     *Commit
	Repository *Repository
	Monitor    *Monitor
}

func New(dep *dep.Dependencies) *Logic {
	once.Do(func() {
		logic = &Logic{
			Commit:     NewCommitLogic(dep.DAL.CommitDAL),
			Repository: NewRepositoryLogic(dep),
			Monitor:    NewMonitorLogic(dep),
		}
	})

	return logic
}
