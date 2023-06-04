package customerServices

import "github.com/nurmeden/PaymentGateway/customer-service/api/customer/repositories"

type CustomerService interface {
}

type customerService struct {
	customerRepository *repositories.CustomerRepository
}

func NewCustomerService(customerRepository *repositories.CustomerRepository) *customerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
