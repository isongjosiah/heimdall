package dal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"github.com/uptrace/bun"
	"heimdall/internal/dal/crudder"
	"heimdall/internal/dal/model"
	"net/url"
	"time"
)

//counterfeiter:generate . ICommitDAL
type ICommitDAL interface {
	TopContributors(ctx context.Context, authorCount, page int) (model.TopContributor, error)
	CommitsByRepoId(ctx context.Context, repoId string, queryValue url.Values) (model.CommitList, error)
	RepoLastCommitDate(ctx context.Context, repoId string) (time.Time, error)
	AddCommits(ctx context.Context, commits []model.Commit) error
	DeleteRepoCommits(ctx context.Context, repoId string) error
}

type SQLCommitDAL struct {
	Db *bun.DB
}

func NewSQLCommitDAL(db *bun.DB) SQLCommitDAL {
	return SQLCommitDAL{
		Db: db,
	}
}

func (scd SQLCommitDAL) TopContributors(ctx context.Context, authorCount, page int) (model.TopContributor, error) {

	var contributors []model.ContributorDetail

	query := `
			SELECT author_name,
				   author_email,
				   COUNT(*) AS commit_count
			FROM (SELECT author ->> 'name'  AS author_name,
						 author ->> 'email' AS author_email
				  FROM commits) AS derived_table
			GROUP BY author_name,
					 author_email
			ORDER BY commit_count DESC LIMIT ?;
             `
	if err := scd.Db.NewRaw(query, authorCount).Scan(context.Background(), &contributors); err != nil {
		return model.TopContributor{}, err
	}

	response := model.TopContributor{
		Contributors: contributors,
	}

	return response, nil

}

func (scd SQLCommitDAL) CommitsByRepoId(ctx context.Context, repoId string, queryValue url.Values) (model.CommitList, error) {

	var commits []model.Commit
	crud := crudder.DefaultCrudder(&commits, scd.Db, queryValue)
	crud.Filter.Exact["repo_id"] = repoId
	err := crud.Fetch()

	response := model.CommitList{
		Commits:        commits,
		PaginationMeta: *crud.Paginator,
	}

	return response, err
}

func (scd SQLCommitDAL) RepoLastCommitDate(ctx context.Context, repoId string) (time.Time, error) {

	var commit model.Commit
	crud := crudder.DefaultCrudder(&commit, scd.Db)
	crud.Filter.Sorter.Desc = []string{"commit_date"}
	err := crud.Fetch()

	return commit.Date, err
}

func (scd SQLCommitDAL) AddCommits(ctx context.Context, commits []model.Commit) error {

	crud := crudder.DefaultCrudder(&commits, scd.Db)
	_, err := crud.Insert()
	return err

}

func (scd SQLCommitDAL) DeleteRepoCommits(ctx context.Context, repoId string) error {

	// extract transaction from context
	tx := ctx.Value("sql-tx").(bun.IDB)
	if tx == nil {
		tx = scd.Db
	}

	crud := crudder.DefaultCrudder(&model.Commit{}, scd.Db)
	crud.Filter.Exact["repo_id"] = repoId
	_, err := crud.Delete()
	return err

}
