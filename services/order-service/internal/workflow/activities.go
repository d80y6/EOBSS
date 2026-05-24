package workflow

import (
	"context"
	"fmt"
	"os"
	"net/http"
	"bytes"
	"encoding/json"
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
	logger.Info("Provisioning service via provisioning-service", zap.String("ServiceOrderID", serviceOrderID))

	provSvcURL := os.Getenv("PROVISIONING_SERVICE_URL")
	if provSvcURL == "" {
		provSvcURL = "http://provisioning-service:8080"
	}

	reqBody, _ := json.Marshal(map[string]string{
		"serviceOrderId": serviceOrderID,
		"action": "activate",
	})

	resp, err := http.Post(provSvcURL+"/provision", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("provisioning service returned status: %d", resp.StatusCode)
	}

	return nil
}

func (a *Activities) ActivateBillingActivity(ctx context.Context, order domain.ProductOrder) error {
	logger.Info("Activating billing for order", zap.String("OrderID", order.ID))

	billingSvcURL := os.Getenv("BILLING_SERVICE_URL")
	if billingSvcURL == "" {
		billingSvcURL = "http://billing-service:8080"
	}

	reqBody, _ := json.Marshal(order)
	resp, err := http.Post(billingSvcURL+"/billing/activate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("billing service returned status: %d", resp.StatusCode)
	}

	return nil
}
