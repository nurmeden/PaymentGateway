package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/services"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type paymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *paymentHandler {
	return &paymentHandler{
		paymentService: paymentService,
	}
}

func (h *paymentHandler) CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.paymentService.CreatePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, payment)

}

// func (h *PaymentHandler) GetPaymentByID(w http.ResponseWriter, r *http.Request) {}

// func (h *PaymentHandler) UpdatePayment(w http.ResponseWriter, r *http.Request) {}

// func (h *PaymentHandler) DeletePayment(w http.ResponseWriter, r *http.Request) {}
