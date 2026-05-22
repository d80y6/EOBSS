package workflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

type ProvisioningRequest struct {
	OrderID    string
	ServiceID  string
	TargetType string // GPON, VoIP, 5G
	Params     map[string]string
}

// ProvisioningSagaWorkflow implements the Saga pattern for multi-step service activation
func ProvisioningSagaWorkflow(ctx workflow.Context, req ProvisioningRequest) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var compensations []string

	// Step 1: Allocate IP Address
	var ip string
	err := workflow.ExecuteActivity(ctx, "AllocateIPActivity", req).Get(ctx, &ip)
	if err != nil {
		return "Failed", m.compensate(ctx, compensations)
	}
	compensations = append(compensations, "ReleaseIPActivity")

	// Step 2: Configure Network Element (Radius/Router)
	err = workflow.ExecuteActivity(ctx, "ConfigureNEActivity", req, ip).Get(ctx, nil)
	if err != nil {
		return "Failed", m.compensate(ctx, compensations)
	}
	compensations = append(compensations, "DeconfigureNEActivity")

	// Step 3: Activate Service in HSS/Core
	err = workflow.ExecuteActivity(ctx, "ActivateCoreActivity", req).Get(ctx, nil)
	if err != nil {
		return "Failed", m.compensate(ctx, compensations)
	}

	return "Activated", nil
}

func (m *ProvisioningSagaWorkflow) compensate(ctx workflow.Context, activities []string) error {
	// Execute compensations in reverse order
	for i := len(activities) - 1; i >= 0; i-- {
		_ = workflow.ExecuteActivity(ctx, activities[i]).Get(ctx, nil)
	}
	return nil
}
