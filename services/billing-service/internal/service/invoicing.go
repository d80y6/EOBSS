package service

import (
	"context"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
	"time"
)

type InvoicingService struct {
	// clickhouseRepo ClickHouseRepository
}

func (s *InvoicingService) GenerateInvoice(ctx context.Context, accountID string, period string) (*domain.Invoice, error) {
	// 1. Fetch all rated CDRs for the account and period from ClickHouse
	// cdrItems, err := s.clickhouseRepo.GetRatedCDRs(ctx, accountID, period)

	// 2. Aggregate into Invoice Items
	invoice := &domain.Invoice{
		ID:             "INV-" + accountID + "-" + period,
		BillingAccount: domain.AccountRef{ID: accountID},
		InvoiceDate:    time.Now(),
		Status:         "Draft",
		TotalAmount:    domain.Money{Amount: 0, Currency: "USD"},
	}

	// Mock aggregation
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
