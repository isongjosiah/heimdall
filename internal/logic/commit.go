package logic

import (
	"context"
	"heimdall/internal/dal"
	"heimdall/internal/dal/model"
	"heimdall/internal/value"
)

type Commit struct {
	DAL dal.ICommitDAL
}

func NewCommitLogic(dal dal.ICommitDAL) *Commit {
	return &Commit{
		DAL: dal,
	}
}

/*
TopContributors retrieves the top authorCount contributors
*/
func (c *Commit) TopContributors(ctx context.Context, authorCount, page int) (model.TopContributor, string, string, error) {

	topContributors, err := c.DAL.TopContributors(ctx, authorCount, page)
	if err != nil {
		return topContributors, value.Error, "Unable to retrieve top contributors. Please try again", err
	}

	return topContributors, value.Success, "Top contributors retrieved", nil

}
