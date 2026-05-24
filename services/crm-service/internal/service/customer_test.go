package service

import (
	"context"
	"testing"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
)

func TestCreateCustomer(t *testing.T) {
	logger.InitLogger("test", "info")
	repo := NewMockCustomerRepo()
	svc := NewCustomerService(repo, nil)

	customer := &domain.Customer{
		ID:   "123",
		Name: "John Doe",
	}

	err := svc.CreateCustomer(context.Background(), customer)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	if customer.Status != "Active" {
		t.Errorf("expected status Active, got %s", customer.Status)
	}
}
