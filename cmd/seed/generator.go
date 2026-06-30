package main

import (
	"fmt"
	"math/rand"

	"gear-priority-api/internal/dto"
	"gear-priority-api/internal/utils"

	"github.com/google/uuid"
)

var categories = []string{
	"engine",
	"brakes",
	"suspension",
	"transmission",
	"electrical",
	"cooling",
}

var names = []string{
	"Filtro de Óleo",
	"Pastilha de Freio",
	"Disco de Freio",
	"Radiador",
	"Amortecedor",
	"Correia Dentada",
	"Coxim do Motor",
	"Bomba de Combustível",
	"Filtro de Ar",
	"Bateria",
	"Alternador",
	"Motor de Partida",
	"Velas",
	"Pivô",
	"Rolamento",
}

/*
The following gears intentionally represent edge cases.

They allow us to validate the robustness of the
priority calculation algorithm.

Examples:
- Negative stock
- Zero sales
- Extremely long lead time
- High criticality
*/
var edgeCases = []dto.Gear{
	{
		ID:                uuid.New(),
		Name:              "EDGE - Negative Stock",
		Category:          "engine",
		CurrentStock:      -10,
		MinimumStock:      20,
		AverageDailySales: 6,
		LeadTimeInDays:    15,
		UnitCost:          120,
		CriticalityLevel:  5,
	},
	{
		ID:                uuid.New(),
		Name:              "EDGE - Zero Sales",
		Category:          "engine",
		CurrentStock:      50,
		MinimumStock:      20,
		AverageDailySales: 0,
		LeadTimeInDays:    30,
		UnitCost:          50,
		CriticalityLevel:  5,
	},
	{
		ID:                uuid.New(),
		Name:              "EDGE - Long Lead Time",
		Category:          "electrical",
		CurrentStock:      40,
		MinimumStock:      20,
		AverageDailySales: 2,
		LeadTimeInDays:    120,
		UnitCost:          300,
		CriticalityLevel:  4,
	},
}

func generateGear(
	r *rand.Rand,
	index int,
	criticalLimit int,
	mediumLimit int,
) dto.Gear {

	/*
		Instead of generating completely random values,
		we generate three inventory profiles.

		This creates a much more realistic database and
		allows the priority algorithm to be properly tested.

		20% -> Critical
		40% -> Medium
		40% -> Healthy
	*/

	switch {

	case index < criticalLimit:
		return generateCriticalGear(r, index)

	case index < mediumLimit:
		return generateMediumGear(r, index)

	default:
		return generateHealthyGear(r, index)
	}
}

func generateCriticalGear(r *rand.Rand, index int) dto.Gear {

	return dto.Gear{
		ID: uuid.New(),

		Name: fmt.Sprintf(
			"%s %d",
			names[r.Intn(len(names))],
			index,
		),

		Category: categories[r.Intn(len(categories))],

		CurrentStock: r.Intn(10),

		MinimumStock: 30 + r.Intn(20),

		AverageDailySales: float64(
			8 + r.Intn(12),
		),

		LeadTimeInDays: 20 + r.Intn(20),

		UnitCost: utils.RoundMoney(50 + r.Float64()*500),

		CriticalityLevel: 5,
	}
}

func generateMediumGear(r *rand.Rand, index int) dto.Gear {

	return dto.Gear{
		ID: uuid.New(),

		Name: fmt.Sprintf(
			"%s %d",
			names[r.Intn(len(names))],
			index,
		),

		Category: categories[r.Intn(len(categories))],

		CurrentStock: 15 + r.Intn(30),

		MinimumStock: 20 + r.Intn(15),

		AverageDailySales: float64(
			3 + r.Intn(8),
		),

		LeadTimeInDays: 10 + r.Intn(10),

		UnitCost: utils.RoundMoney(30 + r.Float64()*300),

		CriticalityLevel: 3 + r.Intn(2),
	}
}

func generateHealthyGear(r *rand.Rand, index int) dto.Gear {

	return dto.Gear{
		ID: uuid.New(),

		Name: fmt.Sprintf(
			"%s %d",
			names[r.Intn(len(names))],
			index,
		),

		Category: categories[r.Intn(len(categories))],

		CurrentStock: 70 + r.Intn(100),

		MinimumStock: 10 + r.Intn(20),

		AverageDailySales: float64(
			1 + r.Intn(4),
		),

		LeadTimeInDays: 1 + r.Intn(5),

		UnitCost: utils.RoundMoney(10 + r.Float64()*150),

		CriticalityLevel: 1 + r.Intn(3),
	}
}
