package service

import (
	"context"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/resilience"
)

type OrderService struct {
	cb *resilience.CircuitBreaker
}

func NewOrderService() *OrderService {
	return &OrderService{
		cb: resilience.NewCircuitBreaker("order-service"),
	}
}

func (s *OrderService) ProcessOrder(ctx context.Context) error {
	logger.Log.Info("Processing order with resilience")
	_, err := s.cb.Execute(func() (interface{}, error) {
		// Call to external service or DB
		return nil, nil
	})
	return err
}
