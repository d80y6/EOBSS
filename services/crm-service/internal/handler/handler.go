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
		tmfErr := errors.NewBadRequestError("Invalid request body")
		return c.JSON(tmfErr.Status, tmfErr)
	}

	if err := h.service.CreateCustomer(c.Request().Context(), &customer); err != nil {
		tmfErr := errors.NewInternalError(err.Error())
		return c.JSON(tmfErr.Status, tmfErr)
	}

	return c.JSON(http.StatusCreated, customer)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	customer, err := h.service.GetCustomer(c.Request().Context(), id)
	if err != nil {
		tmfErr := errors.NewNotFoundError("Customer not found")
		return c.JSON(tmfErr.Status, tmfErr)
	}

	return c.JSON(http.StatusOK, customer)
}
