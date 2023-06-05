package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	customerServices "github.com/nurmeden/PaymentGateway/customer-service/api/customer/services"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type customerHandler struct {
	customerService customerServices.CustomerService
}

func NewCustomerHandler(customerServices customerServices.CustomerService) customerHandler {
	return customerHandler{
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

func (h *customerHandler) GetCustomerByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}

	customer, err := h.customerService.GetCustomerByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

func (h *customerHandler) UpdateCustomer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}

	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	customer.ID = id

	err = h.customerService.UpdateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customer)

}

func (h *customerHandler) DeleteCustomer(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid customer ID")
	}

	err = h.customerService.DeleteCustomer(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Customer deleted successfully")
}
