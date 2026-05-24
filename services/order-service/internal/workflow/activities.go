package workflow

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/logger"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/resilience"
	"github.com/telcoflow/telcoflow/services/order-service/internal/domain"
	"go.uber.org/zap"
)

type Activities struct {
	provisionCB *resilience.CircuitBreaker
	billingCB   *resilience.CircuitBreaker
}

func NewActivities() *Activities {
	return &Activities{
		provisionCB: resilience.NewCircuitBreaker("provisioning-service"),
		billingCB:   resilience.NewCircuitBreaker("billing-service"),
	}
}

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
		soID := fmt.Sprintf("SO-%s-%d", order.ID, i)
		serviceOrders = append(serviceOrders, soID)
	}

	return serviceOrders, nil
}

func (a *Activities) ProvisionServiceActivity(ctx context.Context, serviceOrderID string) error {
	logger.Info("Provisioning service via provisioning-service", zap.String("ServiceOrderID", serviceOrderID))

	_, err := a.provisionCB.Execute(func() (interface{}, error) {
		provSvcURL := os.Getenv("PROVISIONING_SERVICE_URL")
		if provSvcURL == "" {
			provSvcURL = "http://provisioning-service:8080"
		}

		reqBody, _ := json.Marshal(map[string]string{
			"serviceOrderId": serviceOrderID,
			"action":         "activate",
		})

		resp, err := http.Post(provSvcURL+"/provision", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
			return nil, fmt.Errorf("provisioning service returned status: %d", resp.StatusCode)
		}
		return nil, nil
	})

	return err
}

func (a *Activities) ActivateBillingActivity(ctx context.Context, order domain.ProductOrder) error {
	logger.Info("Activating billing for order", zap.String("OrderID", order.ID))

	_, err := a.billingCB.Execute(func() (interface{}, error) {
		billingSvcURL := os.Getenv("BILLING_SERVICE_URL")
		if billingSvcURL == "" {
			billingSvcURL = "http://billing-service:8080"
		}

		reqBody, _ := json.Marshal(order)
		resp, err := http.Post(billingSvcURL+"/billing/activate", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("billing service returned status: %d", resp.StatusCode)
		}
		return nil, nil
	})

	return err
}
