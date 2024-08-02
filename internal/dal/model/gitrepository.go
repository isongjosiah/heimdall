package model

import "time"

type GitRepository struct {
	Id              string    `json:"id" bun:"id"`
	Name            string    `json:"name" bun:"name"`
	Owner           string    `json:"owner" bun:"owner"`
	Description     string    `json:"description" bun:"description"`
	URL             string    `json:"url" bun:"url"`
	Language        string    `json:"language" bun:"language"`
	ForkCount       int       `json:"fork_count" bun:"fork_count"`
	StarsCount      int       `json:"stars_count" bun:"stars_count"`
	OpenIssueCount  int       `json:"open_issue_count" bun:"open_issue_count"`
	WatchersCount   int       `json:"watchers_count" bun:"watchers_count"`
	PullFrom        time.Time `json:"pull_from" bun:"pull_from"`
	InitialPullDone bool      `json:"-" bun:"initial_pull_done"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at"`
	AddedAt         time.Time `json:"added_at" bun:"added_at"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at"`
}

type AddRepositoryInput struct {
	RepoName        string    `json:"repo_name"`
	RepoId          string    `json:"repo_id"`
	RepoOwner       string    `json:"repo_owner"`
	FetchCommitFrom time.Time `json:"fetch_commit_from"`
}

type ResetCollectionInput struct {
	RepoName  string    `json:"repo_name"`
	StartFrom time.Time `json:"start_from"`
}
