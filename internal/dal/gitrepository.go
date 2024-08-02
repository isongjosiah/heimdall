package dal

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"heimdall/internal/dal/crudder"
	"heimdall/internal/dal/model"
	"heimdall/internal/value"
)

//counterfeiter:generate . IGitRepositoryDAL
type IGitRepositoryDAL interface {
	RepoByName(ctx context.Context, repoName string) (model.GitRepository, string, error)
	RepoExists(ctx context.Context, repoName string) (bool, error)
	ListRepoCursor(ctx context.Context, cursor string, limit int) ([]model.GitRepository, error)
	AddRepo(ctx context.Context, repo model.GitRepository) error
	UpdateRepo(ctx context.Context, id string, updates map[string]any) error
}

type SQLGitRepositoryDAL struct {
	DB bun.IDB
}

func NewSQLGitRepository(db *bun.DB) *SQLGitRepositoryDAL {
	return &SQLGitRepositoryDAL{
		DB: db,
	}
}

func (rd SQLGitRepositoryDAL) RepoExists(ctx context.Context, repoName string) (bool, error) {

	var repo model.GitRepository
	crud := crudder.DefaultCrudder(&repo, rd.DB)
	crud.Filter.Exact["name"] = repoName
	return crud.Exists()

}

func (rd SQLGitRepositoryDAL) RepoByName(ctx context.Context, repoName string) (model.GitRepository, string, error) {

	var repo model.GitRepository
	crud := crudder.DefaultCrudder(&repo, rd.DB)
	crud.Filter.Exact["name"] = repoName
	if err := crud.Fetch(); err != nil {
		status := value.Error
		if errors.Is(err, sql.ErrNoRows) {
			status = value.NotFound
		}
		return repo, status, err
	}

	return repo, value.Success, nil
}

func (rd SQLGitRepositoryDAL) AddRepo(ctx context.Context, repo model.GitRepository) error {

	// extract transaction from context
	tx := ctx.Value("sql-tx").(bun.IDB)
	if tx == nil {
		tx = rd.DB
	}

	crud := crudder.DefaultCrudder(&repo, tx)
	_, err := crud.Insert()
	return err

}

func (rd SQLGitRepositoryDAL) ListRepoCursor(ctx context.Context, cursor string, limit int) ([]model.GitRepository, error) {

	var repos []model.GitRepository
	crud := crudder.DefaultCrudder(&repos, rd.DB)
	crud.Filter.Exact["initial_pull_done"] = true
	crud.Filter.RawWhere = []string{"id > " + cursor}
	crud.Paginator = &crudder.PaginationMeta{
		Page:    1,
		PerPage: 10,
	}
	err := crud.Fetch()
	return repos, err

}

func (rd SQLGitRepositoryDAL) UpdateRepo(ctx context.Context, id string, updates map[string]any) error {

	crud := crudder.DefaultCrudder(&model.GitRepository{}, rd.DB)
	crud.Filter.Exact["id"] = id
	crud.Setter.Default = updates
	_, err := crud.Update()
	return err

}
