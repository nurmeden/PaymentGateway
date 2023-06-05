package customerServices

import (
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/repositories"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByID(id int64) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
	DeleteCustomer(id int64) error
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

func (s *customerService) GetCustomerByID(id int64) (*models.Customer, error) {
	customer, err := s.customerRepository.GetCustomerByID(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerService) UpdateCustomer(customer *models.Customer) error {
	err := s.customerRepository.UpdateCustomer(customer)
	if err != nil {
		return err
	}

	return nil
}

func (s *customerService) DeleteCustomer(id int64) error {
	err := s.customerRepository.DeleteCustomer(id)
	if err != nil {
		return err
	}

	return nil
}
