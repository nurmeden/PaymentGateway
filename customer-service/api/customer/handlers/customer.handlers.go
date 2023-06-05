package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	customerServices "github.com/nurmeden/PaymentGateway/customer-service/api/customer/services"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type customerHandler struct {
	customerService customerServices.CustomerService
}

func NewCustomerHandler(customerServices customerServices.CustomerService) *customerHandler {
	return &customerHandler{
		customerService: customerServices,
	}
}

func (h *customerHandler) CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.customerService.CreateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customer)
}
