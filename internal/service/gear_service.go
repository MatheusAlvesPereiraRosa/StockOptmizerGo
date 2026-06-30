package service

import (
	"context"
	"math"

	"gear-priority-api/internal/domain"
	"gear-priority-api/internal/dto"
	"gear-priority-api/internal/utils"
)

type GearService struct {
	repository domain.GearRepository
}

const (
	defaultPage  = 1
	defaultLimit = 20
	maxLimit     = 100
)

/* Make the repository creation independent from the service, so that we can easily swap out the repository implementation if needed. */
func NewGearService(repository domain.GearRepository) *GearService {
	return &GearService{
		repository: repository,
	}
}

func (s *GearService) Create(
	ctx context.Context,
	gear *dto.Gear,
) error {
	return s.repository.Create(ctx, gear)
}

func (s *GearService) FindAll(
	ctx context.Context,
) ([]dto.Gear, error) {
	return s.repository.FindAll(ctx)
}

func (s *GearService) FindByID(
	ctx context.Context,
	id string,
) (*dto.Gear, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *GearService) FindPaginated(
	ctx context.Context,
	page, limit int,
) (*dto.PaginatedGearsResponse, error) {
	page, limit = utils.NormalizePagination(
		page,
		limit,
		defaultPage,
		defaultLimit,
		maxLimit,
	)

	gears, totalItems, err := s.repository.FindPaginated(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &dto.PaginatedGearsResponse{
		Data: gears,
		Meta: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *GearService) FindByCategory(
	ctx context.Context,
	category string,
) ([]dto.Gear, error) {
	return s.repository.FindByCategory(ctx, category)
}

func (s *GearService) FindByCategoryPaginated(
	ctx context.Context,
	category string,
	page, limit int,
) (*dto.PaginatedGearsResponse, error) {
	page, limit = utils.NormalizePagination(
		page,
		limit,
		defaultPage,
		defaultLimit,
		maxLimit,
	)

	gears, totalItems, err := s.repository.FindByCategoryPaginated(
		ctx,
		category,
		page,
		limit,
	)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	return &dto.PaginatedGearsResponse{
		Data: gears,
		Meta: dto.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *GearService) Update(
	ctx context.Context,
	gear *dto.Gear,
) error {
	return s.repository.Update(ctx, gear)
}

func (s *GearService) Delete(
	ctx context.Context,
	id string,
) error {
	return s.repository.Delete(ctx, id)
}
