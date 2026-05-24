package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/services/inventory-service/internal/service"
	"github.com/telcoflow/telcoflow/services/inventory-service/internal/domain"
)

type InventoryHandler struct {
	svc *service.InventoryService
}

func NewInventoryHandler(svc *service.InventoryService) *InventoryHandler {
	return &InventoryHandler{svc: svc}
}

func (h *InventoryHandler) GetResource(c echo.Context) error {
	id := c.Param("id")
	res, err := h.svc.GetResource(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "resource not found"})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *InventoryHandler) CreateResource(c echo.Context) error {
	var res domain.Resource
	if err := c.Bind(&res); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid resource structure"})
	}
	if err := h.svc.CreateResource(c.Request().Context(), &res); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, res)
}
