package service

import (
	"context"

	"soccer-api/internal/entity"
	"soccer-api/internal/team"
)

type teamService struct {
	repo team.Team
}

// New will return teamService
func New(repo team.Team) team.Team {
	return teamService{
		repo: repo,
	}
}

func (ts teamService) Fetch(ctx context.Context, filter entity.QueryFilter) ([]entity.Team, string, error) {
	return ts.repo.Fetch(ctx, filter)
}

func (ts teamService) Get(ctx context.Context, ID int64) (entity.Team, error) {
	return ts.repo.Get(ctx, ID)
}

func (ts teamService) Insert(ctx context.Context, t entity.Team) (entity.Team, error) {
	return ts.repo.Insert(ctx, t)
}
