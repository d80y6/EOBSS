package service

import (
	"context"
	"fmt"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
)

type mockCustomerRepo struct {
	customers map[string]*domain.Customer
}

func NewMockCustomerRepo() CustomerRepository {
	return &mockCustomerRepo{customers: make(map[string]*domain.Customer)}
}

func (r *mockCustomerRepo) Create(ctx context.Context, customer *domain.Customer) error {
	r.customers[customer.ID] = customer
	return nil
}

func (r *mockCustomerRepo) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	customer, ok := r.customers[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return customer, nil
}

func (r *mockCustomerRepo) Update(ctx context.Context, customer *domain.Customer) error {
	r.customers[customer.ID] = customer
	return nil
}
