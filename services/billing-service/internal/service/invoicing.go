package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/rating"
)

type InvoicingService struct {
	ratingEngine *rating.RatingEngine
}

func NewInvoicingService(engine *rating.RatingEngine) *InvoicingService {
	return &InvoicingService{ratingEngine: engine}
}

func (s *InvoicingService) GenerateInvoice(ctx context.Context, customerID string, usageRecords []domain.UsageRecord) (*domain.Invoice, error) {
	totalAmount := 0.0
	for _, record := range usageRecords {
		// Ensure record is rated
		if record.RatedAmount == 0 {
			_ = s.ratingEngine.RateUsage(ctx, &record)
		}
		totalAmount += record.RatedAmount
	}

	invoice := &domain.Invoice{
		ID:         "INV-" + customerID,
		CustomerID: customerID,
		Amount:     totalAmount,
		Currency:   "USD",
		Status:     "Generated",
	}
	return invoice, nil
}

func (s *InvoicingService) ProcessUsage(ctx context.Context, record *domain.UsageRecord) error {
	return s.ratingEngine.RateUsage(ctx, record)
}
