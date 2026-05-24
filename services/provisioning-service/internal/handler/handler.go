package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/errors"
	"github.com/telcoflow/telcoflow/services/provisioning-service/internal/service"
)

type ProvisioningHandler struct {
	svc *service.ProvisioningService
}

func NewProvisioningHandler(svc *service.ProvisioningService) *ProvisioningHandler {
	return &ProvisioningHandler{svc: svc}
}

func (h *ProvisioningHandler) Provision(c echo.Context) error {
	var req struct {
		ServiceOrderID string `json:"serviceOrderId"`
		Action         string `json:"action"`
		Type           string `json:"type"` // radius, voip, 5g
		Username       string `json:"username"`
		Password       string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid request body"))
	}

	var err error
	switch req.Type {
	case "radius":
		err = h.svc.ProvisionRadius(c.Request().Context(), req.Username, req.Password)
	case "voip":
		err = h.svc.ProvisionVoIP(c.Request().Context(), req.Username, "telcoflow.local", req.Password)
	case "5g":
		err = h.svc.Provision5G(c.Request().Context(), req.Username, req.Username) // Mock IMSI/MSISDN
	default:
		// Default to success for generic activation
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "provisioned", "serviceOrderId": req.ServiceOrderID})
}
