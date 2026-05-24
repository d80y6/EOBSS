package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/errors"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/billing-service/internal/service"
)

type BillingHandler struct {
	svc *service.InvoicingService
}

func NewBillingHandler(svc *service.InvoicingService) *BillingHandler {
	return &BillingHandler{svc: svc}
}

func (h *BillingHandler) GenerateInvoice(c echo.Context) error {
	customerID := c.Param("customerId")
	// In a real system, we'd fetch usage records from the repository first
	invoice, err := h.svc.GenerateInvoice(c.Request().Context(), customerID, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}
	return c.JSON(http.StatusOK, invoice)
}

func (h *BillingHandler) ProcessUsage(c echo.Context) error {
	var record domain.UsageRecord
	if err := c.Bind(&record); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid request body"))
	}

	if err := h.svc.ProcessUsage(c.Request().Context(), &record); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusOK, record)
}

func (h *BillingHandler) ActivateBilling(c echo.Context) error {
	// Placeholder for TMF622 integration
	return c.JSON(http.StatusOK, map[string]string{"status": "activated"})
}
