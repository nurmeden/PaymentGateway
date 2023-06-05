package repositories

import "database/sql"

type CustomerRepository interface {
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
	return &customerRepository{
		db: db,
	}
}
