package service

import (
	"context"
	"gear-priority-api/internal/domain"
	"gear-priority-api/internal/dto"
)

type RestockService struct {
	repository domain.GearRepository
}

func NewRestockService(repository domain.GearRepository) *RestockService {
	return &RestockService{
		repository: repository,
	}
}

func (s *RestockService) GetPriorities(
	ctx context.Context,
) (*dto.RestockPriorityResponse, error) {
	/* Get all the gears form the repository to calculate the priority*/
	gears, err := s.repository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	/* Calculate the priorities based on the gears retrieved from the repository */
	priorities := calculatePriorities(gears)

	return &dto.RestockPriorityResponse{
		Priorities: priorities,
	}, nil
}
