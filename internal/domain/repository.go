package domain

import (
	"context"

	"gear-priority-api/internal/dto"
)

type GearRepository interface {
	Create(ctx context.Context, gear *dto.Gear) error
	Update(ctx context.Context, gear *dto.Gear) error
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context) error

	FindByID(ctx context.Context, id string) (*dto.Gear, error)
	FindAll(ctx context.Context) ([]dto.Gear, error)
	FindByCategory(ctx context.Context, category string) ([]dto.Gear, error)

	FindPaginated(ctx context.Context, page, limit int) ([]dto.Gear, int64, error)
	FindByCategoryPaginated(
		ctx context.Context,
		category string,
		page, limit int,
	) ([]dto.Gear, int64, error)
}
