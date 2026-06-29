package service

import "gear-priority-api/internal/dto"

func toDTO(
	candidates []dto.PriorityCandidate,
) []dto.RestockPriority {
	priorities := make(
		[]dto.RestockPriority,
		0,
		len(candidates),
	)

	for _, candidate := range candidates {
		priorities = append(
			priorities,
			candidate.RestockPriority,
		)
	}

	return priorities
}