package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/errors"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/catalog-service/internal/service"
)

type CatalogHandler struct {
	svc *service.CatalogService
}

func NewCatalogHandler(svc *service.CatalogService) *CatalogHandler {
	return &CatalogHandler{svc: svc}
}

func (h *CatalogHandler) CreateOffering(c echo.Context) error {
	var offering domain.ProductOffering
	if err := c.Bind(&offering); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid request body"))
	}

	if err := h.svc.CreateOffering(c.Request().Context(), &offering); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusCreated, offering)
}

func (h *CatalogHandler) GetOffering(c echo.Context) error {
	id := c.Param("id")
	offering, err := h.svc.GetOffering(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.NewNotFoundError("product offering not found"))
	}
	return c.JSON(http.StatusOK, offering)
}

func (h *CatalogHandler) ListOfferings(c echo.Context) error {
	offerings, err := h.svc.ListOfferings(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}
	return c.JSON(http.StatusOK, offerings)
}
