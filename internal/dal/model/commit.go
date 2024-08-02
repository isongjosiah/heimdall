package model

import (
	"heimdall/internal/dal/crudder"
	"time"
)

type CommitAuthor struct {
	Name  string `json:"name" bun:"name"`
	Email string `json:"email" bun:"email"`
}

// Commit defines the commit information persisted
type Commit struct {
	Id        string       `json:"id" bun:"id,pk"`
	RepoId    string       `json:"repo_id" bun:"repo_id,unique:repo_commit"`
	Sha       string       `json:"sha" bun:"sha,unique:repo_commit"`
	Message   string       `json:"message" bun:"message"`
	Author    CommitAuthor `json:"author" bun:"author"`
	Date      time.Time    `json:"commit_date" bun:"commit_date"`
	URL       string       `json:"url" bun:"url"`
	AddedAt   time.Time    `json:"added_at" bun:"added_at"`
	DeletedAt time.Time    `json:"deleted_at" bun:"deleted_at,soft_delete,nullzero"`
}

type ContributorDetail struct {
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
	CommitCount int    `json:"commit_count"`
}

type TopContributor struct {
	Contributors   []ContributorDetail    `json:"contributors"`
	PaginationMeta crudder.PaginationMeta `json:"pagination_meta"`
}

type CommitList struct {
	Commits        []Commit               `json:"commits"`
	PaginationMeta crudder.PaginationMeta `json:"pagination_meta"`
}
