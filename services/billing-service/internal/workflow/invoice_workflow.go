package workflow

import (
	"time"
	"go.temporal.io/sdk/workflow"
)

type InvoiceRequest struct {
	BillingAccountID string
	BillingPeriod    string
}

// MonthlyInvoiceWorkflow orchestrates the collection of rated usage and generation of PDF invoices
func MonthlyInvoiceWorkflow(ctx workflow.Context, req InvoiceRequest) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// 1. Fetch Rated Usage from ClickHouse
	var totalAmount float64
	err := workflow.ExecuteActivity(ctx, "AggregateUsageActivity", req).Get(ctx, &totalAmount)
	if err != nil {
		return "Failed", err
	}

	// 2. Apply Taxes & Credits
	var finalAmount float64
	err = workflow.ExecuteActivity(ctx, "CalculateTaxesActivity", totalAmount).Get(ctx, &finalAmount)
	if err != nil {
		return "Failed", err
	}

	// 3. Generate Invoice PDF
	var pdfURL string
	err = workflow.ExecuteActivity(ctx, "GeneratePDFActivity", req, finalAmount).Get(ctx, &pdfURL)
	if err != nil {
		return "Failed", err
	}

	// 4. Notify Customer
	err = workflow.ExecuteActivity(ctx, "NotifyInvoiceActivity", req, pdfURL).Get(ctx, nil)
	if err != nil {
		return "Notification Failed", nil
	}

	return "Invoice Generated", nil
}
