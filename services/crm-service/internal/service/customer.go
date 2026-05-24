package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/kafka"
	"go.uber.org/zap"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *domain.Customer) error
	GetByID(ctx context.Context, id string) (*domain.Customer, error)
	Update(ctx context.Context, customer *domain.Customer) error
}

type CustomerService struct {
	repo     CustomerRepository
	producer *kafka.Producer
}

func NewCustomerService(repo CustomerRepository, producer *kafka.Producer) *CustomerService {
	return &CustomerService{
		repo: repo,
		producer: producer,
	}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *domain.Customer) error {
	if customer.Name == "" {
		return fmt.Errorf("customer name is required")
	}
	customer.Status = "Active"
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	logger.Log.Info("Creating customer", zap.String("name", customer.Name))
	if err := s.repo.Create(ctx, customer); err != nil {
		return err
	}

	if s.producer != nil {
		data, _ := json.Marshal(customer)
		_ = s.producer.PublishEvent(ctx, customer.ID, data)
	}

	return nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, id string) (*domain.Customer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer *domain.Customer) error {
	customer.UpdatedAt = time.Now()
	logger.Log.Info("Updating customer", zap.String("id", customer.ID))
	if err := s.repo.Update(ctx, customer); err != nil {
		return err
	}

	if s.producer != nil {
		data, _ := json.Marshal(customer)
		_ = s.producer.PublishEvent(ctx, customer.ID, data)
	}

	return nil
}
