package domain

import "context"

type Positions struct {
	Id            string `db:"id"`
	DepartementId string `db:"departement_id"`
	Name          string `db:"name"`
	IsActive      bool   `db:"is_active"`
	CreatedAt     string `db:"created_at"`
}

type PositionsRepository interface {
	FindAllPositions(ctx context.Context) (result []Positions, err error)
	FindPositionById(ctx context.Context, positionId string) (result Positions, err error)
	SavePosition(ctx context.Context, position Positions) error
	DeletePosition(ctx context.Context, positionId string) error
}

type PositionsService interface {
	FindAllPositions(ctx context.Context) (result []Positions, err error)
	FindPositionById(ctx context.Context, positionId string) (result Positions, err error)
	CreatePosition(ctx context.Context, position Positions) error
	DeletePosition(ctx context.Context, positionId string) error
}
