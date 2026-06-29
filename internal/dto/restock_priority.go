package dto

import "github.com/google/uuid"

type RestockPriority struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	CurrentStock   int       `json:"currentStock"`
	ProjectedStock float64   `json:"projectedStock"`
	MinimumStock   int       `json:"minimumStock"`
	UrgencyScore   float64   `json:"urgencyScore"`
}

type PriorityCandidate struct {
	RestockPriority

	CriticalityLevel  int     `json:"criticalityLevel"`
	AverageDailySales float64 `json:"averageDailySales"`
}

type RestockPriorityResponse struct {
	Priorities []RestockPriority `json:"priorities"`
}
