package domain

import "github.com/google/uuid"

type Gear struct {
	id                uuid.UUID `json:"id" bson:"_id"`
	name              string    `json:"name" bson:"name"`
	category          string    `json:"category" bson:"category"`
	currentStock      int       `json:"currentStock" bson:"currentStock"`
	minimumStock      int       `json:"minimumStock" bson:"minimumStock"`
	averageDailySales float64   `json:"averageDailySales" bson:"averageDailySales"`
	leadTimeInDays    int       `json:"leadTimeInDays" bson:"leadTimeInDays"`
	unitCost          float64   `json:"unitCost" bson:"unitCost"`
	criticalityLevel  int       `json:"criticalityLevel" bson:"criticalityLevel"`
}
