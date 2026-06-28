package domain

import (
	"context"

	"github.com/google/uuid"
)

type GearRepository interface {
	Create(ctx context.Context, gear *Gear) error
	Update(ctx context.Context, gear *Gear) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*Gear, error)
	FindAll(ctx context.Context) ([]Gear, error)
}
