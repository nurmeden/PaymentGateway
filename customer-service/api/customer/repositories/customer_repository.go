package repositories

import (
	"database/sql"

	"github.com/nurmeden/PaymentGateway/shared/models"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) CreateCustomer(customer *models.Customer) error {

	query := "INSERT INTO customers (name, email, phone) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, customer.Name, customer.Email, customer.Phone)
	if err != nil {
		return err
	}

	return nil
}
