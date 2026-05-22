package rating

import (
	"context"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
)

type PriceTier struct {
	FromQuantity float64
	ToQuantity   float64
	UnitPrice    float64
}

type PricePlan struct {
	ProductID string
	Currency  string
	Tiers     []PriceTier
}

// RatingEngine handles complex usage-based charging
type RatingEngine struct {
	plans map[string]PricePlan
}

func NewRatingEngine() *RatingEngine {
	return &RatingEngine{
		plans: make(map[string]PricePlan),
	}
}

func (e *RatingEngine) RateUsage(ctx context.Context, record *domain.UsageRecord) error {
	// In a real system, we would fetch the plan based on the service/product from a cache or database
	plan, ok := e.plans[record.UsageType]
	if !ok {
		// Default fallback for foundation demo
		rate := 0.01
		record.RatedAmount = record.UsageQuantity * rate
		return nil
	}

	amount := 0.0
	remaining := record.UsageQuantity

	for _, tier := range plan.Tiers {
		if remaining <= 0 {
			break
		}

		tierVolume := tier.ToQuantity - tier.FromQuantity
		if tierVolume <= 0 { // Infinite tier
			amount += remaining * tier.UnitPrice
			break
		}

		if remaining > tierVolume {
			amount += tierVolume * tier.UnitPrice
			remaining -= tierVolume
		} else {
			amount += remaining * tier.UnitPrice
			remaining = 0
		}
	}

	record.RatedAmount = amount
	return nil
}
