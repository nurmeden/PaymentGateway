package repositories

import (
	"database/sql"
	"errors"

	"github.com/nurmeden/PaymentGateway/shared/models"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	// GetPaymentById(id string) (models.Payment, error)
	// UpdatePayment(payment *models.Payment) error
	// DeletePayment(id string) error
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (h *paymentRepository) CreatePayment(payment *models.Payment) error {
	query := `
		INSERT INTO payments (amount, status, method)
		VALUES ($2, $3, $4)
	`
	err := h.db.QueryRow(query, payment.Amount, payment.Status, payment.Method)
	if err != nil {
		return errors.New("error queryRow")
	}

	return nil
}
