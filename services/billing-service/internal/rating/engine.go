package rating

import (
	"context"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
)

// RatingEngine handles usage-based charging
type RatingEngine struct {
	// In a real scenario, this would load price plans from the Product Catalog
}

func (e *RatingEngine) RateUsage(ctx context.Context, record *domain.UsageRecord) error {
	// Simplified rating logic for a carrier-grade foundation
	rate := 0.0
	switch record.UsageType {
	case "Data":
		rate = 0.01 // $0.01 per MB
	case "Voice":
		rate = 0.05 // $0.05 per Minute
	case "SMS":
		rate = 0.02 // $0.02 per SMS
	}

	record.RatedAmount = record.UsageQuantity * rate
	return nil
}
