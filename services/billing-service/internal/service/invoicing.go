package service

import (
	"context"
	"fmt"
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
	var invoiceItems []domain.InvoiceItem

	for i, record := range usageRecords {
		// Real-world logic: rate the usage record before adding to invoice
		_ = s.ratingEngine.RateUsage(ctx, &record)

		totalAmount += record.RatedAmount
		invoiceItems = append(invoiceItems, domain.InvoiceItem{
			ID:          fmt.Sprintf("ITEM-%d", i),
			Description: record.UsageType + " Usage",
			Amount:      domain.Money{Amount: record.RatedAmount, Currency: "USD"},
			ServiceID:   record.ServiceID,
		})
	}

	invoice := &domain.Invoice{
		ID:         "INV-" + customerID,
		CustomerID: customerID,
		Amount:     domain.Money{Amount: totalAmount, Currency: "USD"},
		Status:     "Generated",
		Items:      invoiceItems,
	}
	return invoice, nil
}

func (s *InvoicingService) ProcessUsage(ctx context.Context, record *domain.UsageRecord) error {
	return s.ratingEngine.RateUsage(ctx, record)
}
