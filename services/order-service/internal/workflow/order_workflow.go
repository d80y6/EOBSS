package workflow

import (
	"context"
	"time"

	"github.com/telcoflow/telcoflow/services/order-service/internal/domain"
	"go.temporal.io/sdk/workflow"
)

// ProductOrderWorkflow orchestrates the end-to-end fulfillment of a product order
func ProductOrderWorkflow(ctx workflow.Context, order domain.ProductOrder) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy: &workflow.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Starting Product Order Workflow", "OrderID", order.ID)

	// Phase 1: Validation & Credit Check
	var valid bool
	err := workflow.ExecuteActivity(ctx, "ValidateOrderActivity", order).Get(ctx, &valid)
	if err != nil || !valid {
		return "Failed", err
	}

	// Phase 2: Order Decomposition
	// In a real system, we'd split Product Order into Service Orders
	var serviceOrders []string
	err = workflow.ExecuteActivity(ctx, "DecomposeOrderActivity", order).Get(ctx, &serviceOrders)
	if err != nil {
		return "Failed", err
	}

	// Phase 3: Provisioning (Fan-out)
	for _, soID := range serviceOrders {
		err = workflow.ExecuteActivity(ctx, "ProvisionServiceActivity", soID).Get(ctx, nil)
		if err != nil {
			// Handle partial failure (Saga pattern)
			return "Partially Failed", err
		}
	}

	// Phase 4: Billing Activation
	err = workflow.ExecuteActivity(ctx, "ActivateBillingActivity", order).Get(ctx, nil)
	if err != nil {
		return "Billing Activation Failed", err
	}

	logger.Info("Product Order Workflow Completed", "OrderID", order.ID)
	return "Completed", nil
}
