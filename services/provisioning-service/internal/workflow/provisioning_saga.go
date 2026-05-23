package workflow

import (
	"time"
	"go.temporal.io/sdk/workflow"
)

type ProvisioningRequest struct {
	OrderID    string
	ServiceID  string
	TargetType string
	Params     map[string]string
}

func ProvisioningSagaWorkflow(ctx workflow.Context, req ProvisioningRequest) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var compensations []string

	var ip string
	err := workflow.ExecuteActivity(ctx, "AllocateIPActivity", req).Get(ctx, &ip)
	if err != nil {
		compensate(ctx, compensations)
		return "Failed", err
	}
	compensations = append(compensations, "ReleaseIPActivity")

	err = workflow.ExecuteActivity(ctx, "ConfigureNEActivity", req, ip).Get(ctx, nil)
	if err != nil {
		compensate(ctx, compensations)
		return "Failed", err
	}
	compensations = append(compensations, "DeconfigureNEActivity")

	return "Activated", nil
}

func compensate(ctx workflow.Context, activities []string) {
	for i := len(activities) - 1; i >= 0; i-- {
		_ = workflow.ExecuteActivity(ctx, activities[i]).Get(ctx, nil)
	}
}
