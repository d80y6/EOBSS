package service

import (
	"context"
	"fmt"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *domain.Customer) error
	GetByID(ctx context.Context, id string) (*domain.Customer, error)
	Update(ctx context.Context, customer *domain.Customer) error
}

type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *domain.Customer) error {
	if customer.Name == "" {
		return fmt.Errorf("customer name is required")
	}
	customer.Status = "Active"
	return s.repo.Create(ctx, customer)
}

func (s *CustomerService) GetCustomer(ctx context.Context, id string) (*domain.Customer, error) {
	return s.repo.GetByID(ctx, id)
}
