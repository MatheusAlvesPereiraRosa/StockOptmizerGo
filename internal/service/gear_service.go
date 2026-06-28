package service

import (
	"context"
	"gear-priority-api/internal/domain"

	"github.com/google/uuid"
)

type GearService struct {
	repository domain.GearRepository
}

/* Make the repository creation independent from the service, so that we can easily swap out the repository implementation if needed. */
func NewGearService(repository domain.GearRepository) *GearService {
	return &GearService{
		repository: repository,
	}
}

func (s *GearService) Create(
	ctx context.Context,
	gear *domain.Gear,
) error {
	return s.repository.Create(ctx, gear)
}

func (s *GearService) FindAll(
	ctx context.Context,
) ([]domain.Gear, error) {
	return s.repository.FindAll(ctx)
}

func (s *GearService) FindByID(
	ctx context.Context,
	id uuid.UUID,
) (*domain.Gear, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *GearService) Update(
	ctx context.Context,
	gear *domain.Gear,
) error {
	return s.repository.Update(ctx, gear)
}

func (s *GearService) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return s.repository.Delete(ctx, id)
}
