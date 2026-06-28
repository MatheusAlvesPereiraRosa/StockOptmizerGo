package domain

import (
	"context"
)

type GearRepository interface {
	Create(ctx context.Context, gear *Gear) error
	Update(ctx context.Context, gear *Gear) error
	Delete(ctx context.Context, id string) error

	FindByID(ctx context.Context, id string) (*Gear, error)
	FindAll(ctx context.Context) ([]Gear, error)
}
