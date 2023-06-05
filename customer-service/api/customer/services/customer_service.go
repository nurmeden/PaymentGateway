package customerServices

import (
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/repositories"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) *customerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (s *customerService) CreateCustomer(customer *models.Customer) error {
	err := s.customerRepository.CreateCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}
