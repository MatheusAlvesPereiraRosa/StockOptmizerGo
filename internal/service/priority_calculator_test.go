package service

import (
	"math"
	"testing"

	"gear-priority-api/internal/dto"

	"github.com/google/uuid"
)

func TestCalculatePriority(t *testing.T) {
	tests := []struct {
		name              string
		gear              dto.Gear
		shouldNeedRestock bool
		wantProjected     float64
		wantUrgency       float64
	}{
		{
			name: "gear needs restock",
			gear: dto.Gear{
				ID:                uuid.New(),
				Name:              "Filtro de Óleo X",
				CurrentStock:      15,
				MinimumStock:      20,
				AverageDailySales: 4,
				LeadTimeInDays:    5,
				CriticalityLevel:  3,
			},
			shouldNeedRestock: true,
			wantProjected:     -5,
			wantUrgency:       75,
		},
		{
			name: "gear does not need restock",
			gear: dto.Gear{
				ID:                uuid.New(),
				Name:              "Estoque saudável",
				CurrentStock:      50,
				MinimumStock:      20,
				AverageDailySales: 4,
				LeadTimeInDays:    5,
				CriticalityLevel:  3,
			},
			shouldNeedRestock: false,
		},
		{
			name: "negative current stock increases urgency",
			gear: dto.Gear{
				ID:                uuid.New(),
				Name:              "Estoque negativo",
				CurrentStock:      -10,
				MinimumStock:      20,
				AverageDailySales: 6,
				LeadTimeInDays:    15,
				CriticalityLevel:  5,
			},
			// expected consumption = 6 * 15 = 90
			// projected stock = -10 - 90 = -100
			// shortage = 20 - (-100) = 120
			// urgency = 120 * 5 = 600
			shouldNeedRestock: true,
			wantProjected:     -100,
			wantUrgency:       600,
		},
		{
			name: "zero sales can still require restock",
			gear: dto.Gear{
				ID:                uuid.New(),
				Name:              "Peça sem vendas",
				CurrentStock:      10,
				MinimumStock:      20,
				AverageDailySales: 0,
				LeadTimeInDays:    30,
				CriticalityLevel:  2,
			},
			// projected = 10
			// shortage = 20 - 10 = 10
			// urgency = 10 * 2 = 20
			shouldNeedRestock: true,
			wantProjected:     10,
			wantUrgency:       20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priority, needsRestock := CalculatePriority(&tt.gear)

			if needsRestock != tt.shouldNeedRestock {
				t.Fatalf(
					"needsRestock = %v; want %v",
					needsRestock,
					tt.shouldNeedRestock,
				)
			}

			if !tt.shouldNeedRestock {
				return
			}

			if priority.Name != tt.gear.Name {
				t.Errorf(
					"name = %q; want %q",
					priority.Name,
					tt.gear.Name,
				)
			}

			if !almostEqual(priority.ProjectedStock, tt.wantProjected) {
				t.Errorf(
					"projectedStock = %.2f; want %.2f",
					priority.ProjectedStock,
					tt.wantProjected,
				)
			}

			if !almostEqual(priority.UrgencyScore, tt.wantUrgency) {
				t.Errorf(
					"urgencyScore = %.2f; want %.2f",
					priority.UrgencyScore,
					tt.wantUrgency,
				)
			}
		})
	}
}

func TestSortPriorities(t *testing.T) {
	priorities := []dto.PriorityCandidate{
		{
			RestockPriority: dto.RestockPriority{
				Name:         "Zulu",
				UrgencyScore: 100,
			},
			CriticalityLevel:  3,
			AverageDailySales: 5,
		},
		{
			RestockPriority: dto.RestockPriority{
				Name:         "Bravo",
				UrgencyScore: 100,
			},
			CriticalityLevel:  4,
			AverageDailySales: 1,
		},
		{
			RestockPriority: dto.RestockPriority{
				Name:         "Alpha",
				UrgencyScore: 100,
			},
			CriticalityLevel:  4,
			AverageDailySales: 10,
		},
		{
			RestockPriority: dto.RestockPriority{
				Name:         "Highest Score",
				UrgencyScore: 200,
			},
			CriticalityLevel:  1,
			AverageDailySales: 1,
		},
		{
			RestockPriority: dto.RestockPriority{
				Name:         "Alphabetical First",
				UrgencyScore: 100,
			},
			CriticalityLevel:  3,
			AverageDailySales: 5,
		},
	}

	sortPriorities(priorities)

	wantOrder := []string{
		"Highest Score",      // highest urgency score
		"Alpha",              // same score: highest criticality, then sales
		"Bravo",              // same score and criticality, lower sales
		"Alphabetical First", // same score, criticality and sales as Zulu
		"Zulu",
	}

	for index, wantName := range wantOrder {
		if priorities[index].Name != wantName {
			t.Errorf(
				"position %d = %q; want %q",
				index,
				priorities[index].Name,
				wantName,
			)
		}
	}
}

func TestCalculatePrioritiesFiltersAndSorts(t *testing.T) {
	gears := []dto.Gear{
		{
			ID:                uuid.New(),
			Name:              "Healthy Gear",
			CurrentStock:      100,
			MinimumStock:      20,
			AverageDailySales: 2,
			LeadTimeInDays:    5,
			CriticalityLevel:  2,
		},
		{
			ID:                uuid.New(),
			Name:              "Medium Priority",
			CurrentStock:      10,
			MinimumStock:      20,
			AverageDailySales: 2,
			LeadTimeInDays:    5,
			CriticalityLevel:  2,
		},
		{
			ID:                uuid.New(),
			Name:              "High Priority",
			CurrentStock:      0,
			MinimumStock:      30,
			AverageDailySales: 10,
			LeadTimeInDays:    10,
			CriticalityLevel:  5,
		},
	}

	priorities := calculatePriorities(gears)

	if len(priorities) != 2 {
		t.Fatalf("priorities length = %d; want 2", len(priorities))
	}

	if priorities[0].Name != "High Priority" {
		t.Errorf(
			"first priority = %q; want %q",
			priorities[0].Name,
			"High Priority",
		)
	}

	if priorities[1].Name != "Medium Priority" {
		t.Errorf(
			"second priority = %q; want %q",
			priorities[1].Name,
			"Medium Priority",
		)
	}
}

func almostEqual(got, want float64) bool {
	const tolerance = 0.0001

	return math.Abs(got-want) < tolerance
}
