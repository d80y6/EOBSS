package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/repository"
	"time"
)

type InvoicingService struct {
	repo *repository.ClickHouseCDRRepository
}

func NewInvoicingService(repo *repository.ClickHouseCDRRepository) *InvoicingService {
	return &InvoicingService{repo: repo}
}

func (s *InvoicingService) GenerateInvoice(ctx context.Context, accountID string, period string) (*domain.Invoice, error) {
	// In a real system, we would query the repo here
	// cdrs, err := s.repo.GetCDRsByAccount(ctx, accountID, period)

	invoice := &domain.Invoice{
		ID:             "INV-" + accountID + "-" + period,
		BillingAccount: domain.AccountRef{ID: accountID},
		InvoiceDate:    time.Now(),
		Status:         "Draft",
		TotalAmount:    domain.Money{Amount: 0, Currency: "USD"},
	}

	// Example item representing aggregated usage
	item := domain.InvoiceItem{
		ID:          "ITEM-001",
		Description: "Usage: Data",
		Quantity:    45.5,
		UnitPrice:   domain.Money{Amount: 0.10, Currency: "USD"},
		TotalAmount: domain.Money{Amount: 4.55, Currency: "USD"},
	}
	invoice.Items = append(invoice.Items, item)
	invoice.TotalAmount.Amount += item.TotalAmount.Amount

	return invoice, nil
}
