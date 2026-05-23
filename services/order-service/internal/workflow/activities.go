package workflow

import (
	"context"
	"fmt"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/services/order-service/internal/domain"
	"go.uber.org/zap"
)

type Activities struct{}

func (a *Activities) ValidateOrderActivity(ctx context.Context, order domain.ProductOrder) (bool, error) {
	logger.Info("Validating order", zap.String("OrderID", order.ID))

	if order.Customer.ID == "" {
		return false, fmt.Errorf("missing customer ID")
	}

	if len(order.OrderItems) == 0 {
		return false, fmt.Errorf("order must have at least one item")
	}

	return true, nil
}

func (a *Activities) DecomposeOrderActivity(ctx context.Context, order domain.ProductOrder) ([]string, error) {
	logger.Info("Decomposing order into service orders", zap.String("OrderID", order.ID))

	var serviceOrders []string
	for i := range order.OrderItems {
		// Mock decomposition logic: each product item becomes a service order
		soID := fmt.Sprintf("SO-%s-%d", order.ID, i)
		serviceOrders = append(serviceOrders, soID)
	}

	return serviceOrders, nil
}

func (a *Activities) ProvisionServiceActivity(ctx context.Context, serviceOrderID string) error {
	logger.Info("Provisioning service", zap.String("ServiceOrderID", serviceOrderID))
	// In a real system, this would call the provisioning-service API
	return nil
}

func (a *Activities) ActivateBillingActivity(ctx context.Context, order domain.ProductOrder) error {
	logger.Info("Activating billing for order", zap.String("OrderID", order.ID))
	// In a real system, this would call the billing-service API
	return nil
}
