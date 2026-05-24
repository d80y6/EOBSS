package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/assurance-service/internal/service"
)

type AlarmHandler struct {
	svc *service.AssuranceService
}

func NewAlarmHandler(svc *service.AssuranceService) *AlarmHandler {
	return &AlarmHandler{svc: svc}
}

func (h *AlarmHandler) HandleAlarm(c echo.Context) error {
	var alarm domain.Alarm
	if err := c.Bind(&alarm); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid alarm structure"})
	}

	if err := h.svc.ProcessAlarm(c.Request().Context(), &alarm); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, alarm)
}
