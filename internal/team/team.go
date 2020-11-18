package team

import (
	"context"

	"soccer-api/internal/entity"
)

// Team is interface for team module
//go:generate mockgen -destination=mocks/mock_team.go -package=mocks . Team
type Team interface {
	Fetch(ctx context.Context, filter entity.QueryFilter) ([]entity.Team, string, error)
	Get(ctx context.Context, ID int64) (entity.Team, error)
	Insert(ctx context.Context, t entity.Team) (entity.Team, error)
}
