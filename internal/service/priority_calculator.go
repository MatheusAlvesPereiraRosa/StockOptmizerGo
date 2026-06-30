package service

import (
	"gear-priority-api/internal/dto"

	"math"
	"sort"
)

/* Calculate the priority computed fields based on the gear's attributes. If the projected stock is above the minimum stock, then the gear is not a candidate for restocking and will be ignored. */
func CalculatePriority(
	gear *dto.Gear,
) (dto.PriorityCandidate, bool) {

	expectedConsumption := gear.AverageDailySales *
		float64(gear.LeadTimeInDays)

	println(gear.Name, "Consumo esperado: ", expectedConsumption)

	projectedStock := float64(gear.CurrentStock) -
		expectedConsumption

	println(gear.Name, "Estoque projetado: ", projectedStock, gear.CurrentStock, expectedConsumption)

	if projectedStock > float64(gear.MinimumStock) {
		return dto.PriorityCandidate{}, false
	}

	stockShortage := math.Max(0, float64(gear.MinimumStock)-projectedStock)

	urgencyScore := stockShortage * float64(gear.CriticalityLevel)

	println(gear.Name, "Nivel de urgência: ", urgencyScore)

	return dto.PriorityCandidate{
		RestockPriority: dto.RestockPriority{
			ID:             gear.ID,
			Name:           gear.Name,
			CurrentStock:   gear.CurrentStock,
			ProjectedStock: projectedStock,
			MinimumStock:   gear.MinimumStock,
			UrgencyScore:   urgencyScore,
			StockShortage:  stockShortage,
		},
		CriticalityLevel:  gear.CriticalityLevel,
		AverageDailySales: gear.AverageDailySales,
	}, true
}

/* Sort the priorities based on the urgency score, criticality level, average daily sales, and name. */
func sortPriorities(priorities []dto.PriorityCandidate) {

	sort.Slice(priorities, func(i, j int) bool {

		if priorities[i].UrgencyScore != priorities[j].UrgencyScore {
			return priorities[i].UrgencyScore >
				priorities[j].UrgencyScore
		}

		if priorities[i].CriticalityLevel != priorities[j].CriticalityLevel {
			return priorities[i].CriticalityLevel >
				priorities[j].CriticalityLevel
		}

		if priorities[i].AverageDailySales != priorities[j].AverageDailySales {
			return priorities[i].AverageDailySales >
				priorities[j].AverageDailySales
		}

		return priorities[i].Name < priorities[j].Name
	})
}

func calculatePriorities(
	gears []dto.Gear,
) []dto.RestockPriority {
	priorities := make([]dto.PriorityCandidate, 0)

	/* Iterate over the gears and append the priority candidates */
	for _, gear := range gears {
		priority, ok := CalculatePriority(&gear)

		if ok {
			priorities = append(priorities, priority)
		}
	}

	/* Sort the priorities based on the urgency score, criticality level, average daily sales, and name. */
	sortPriorities(priorities)

	return toDTO(priorities)
}
