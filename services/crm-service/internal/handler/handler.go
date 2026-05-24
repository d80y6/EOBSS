package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/telcoflow/telcoflow/libs/go-common/pkg/errors"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/domain"
	"github.com/telcoflow/telcoflow/services/crm-service/internal/service"
)

type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(service *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid request body"))
	}

	if err := h.service.CreateCustomer(c.Request().Context(), &customer); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	customer, err := h.service.GetCustomer(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.NewNotFoundError("customer not found"))
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid request body"))
	}
	customer.ID = id

	if err := h.service.UpdateCustomer(c.Request().Context(), &customer); err != nil {
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError(err.Error()))
	}

	return c.JSON(http.StatusOK, customer)
}
