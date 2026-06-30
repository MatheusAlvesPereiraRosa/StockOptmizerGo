package dto

import "github.com/google/uuid"

type Gear struct {
	ID                uuid.UUID `json:"id" bson:"_id"`
	Name              string    `json:"name" bson:"name"`
	Category          string    `json:"category" bson:"category"`
	CurrentStock      int       `json:"currentStock" bson:"currentStock"`
	MinimumStock      int       `json:"minimumStock" bson:"minimumStock"`
	AverageDailySales float64   `json:"averageDailySales" bson:"averageDailySales"`
	LeadTimeInDays    int       `json:"leadTimeInDays" bson:"leadTimeInDays"`
	UnitCost          float64   `json:"unitCost" bson:"unitCost"`
	CriticalityLevel  int       `json:"criticalityLevel" bson:"criticalityLevel"`
}
