package handlers

import customerServices "github.com/nurmeden/PaymentGateway/customer-service/api/customer/services"

type customerHandler struct {
	customerService *customerServices.CustomerService
}

func NewCustomerHandler(customerServices customerServices.CustomerService) *customerHandler {
	return &customerHandler{
		customerService: &customerServices,
	}
}
