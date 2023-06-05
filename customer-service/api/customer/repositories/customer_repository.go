package repositories

import (
	"database/sql"
	"errors"

	"github.com/nurmeden/PaymentGateway/shared/models"
)

type CustomerRepository interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomerByID(id int64) (*models.Customer, error)
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

func (r *customerRepository) GetCustomerByID(id int64) (*models.Customer, error) {
	query := `SELECT id, name, email, phone FROM customers WHERE id = $1 `
	row := r.db.QueryRow(query, id)
	var customer models.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	return &customer, nil

}

func (r *customerRepository) UpdateCustomer(customer *models.Customer) error {
	query := `UPDATE customers SET name = $1, email = $2. phone = $3 WHERE id = $4`
	_, err := r.db.Exec(query, customer.Name, customer.Email, customer.Phone, customer.ID)
	if err != nil {
		return err
	}
	return nil
}
