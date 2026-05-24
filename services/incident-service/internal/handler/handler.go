package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/services/incident-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/incident-service/internal/service"
)

type IncidentHandler struct {
	svc *service.IncidentService
}

func NewIncidentHandler(svc *service.IncidentService) *IncidentHandler {
	return &IncidentHandler{svc: svc}
}

func (h *IncidentHandler) CreateTicket(c echo.Context) error {
	var ticket domain.TroubleTicket
	if err := c.Bind(&ticket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := h.svc.CreateTicket(c.Request().Context(), &ticket); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, ticket)
}
