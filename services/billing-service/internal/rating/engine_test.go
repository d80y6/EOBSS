package rating

import (
	"context"
	"testing"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
)

func TestRateUsage(t *testing.T) {
	engine := NewRatingEngine()

	// Add a plan
	engine.plans["Data"] = PricePlan{
		ProductID: "Data",
		Currency:  "USD",
		Tiers: []PriceTier{
			{FromQuantity: 0, ToQuantity: 100, UnitPrice: 0.10},  // 0-100 MB at -bash.10
			{FromQuantity: 100, ToQuantity: 0, UnitPrice: 0.05}, // Above 100 MB at -bash.05
		},
	}

	tests := []struct {
		name     string
		quantity float64
		expected float64
	}{
		{"Low Usage", 50, 5.0},
		{"Exactly Tier", 100, 10.0},
		{"High Usage", 150, 12.5}, // 100*0.10 + 50*0.05 = 10.0 + 2.5 = 12.5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := &domain.UsageRecord{
				UsageType:     "Data",
				UsageQuantity: tt.quantity,
			}
			err := engine.RateUsage(context.Background(), record)
			if err != nil {
				t.Fatalf("RateUsage failed: %v", err)
			}
			if record.RatedAmount != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, record.RatedAmount)
			}
		})
	}
}
