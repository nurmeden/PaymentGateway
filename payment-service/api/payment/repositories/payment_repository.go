package repositories

import (
	"database/sql"
	"errors"

	"github.com/nurmeden/PaymentGateway/shared/models"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByID(id int64) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) (*models.Payment, error)
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
		INSERT INTO payments (id, amount, status, method)
		VALUES ($1, $2, $3, $4)
	`
	_, err := h.db.Exec(query, payment.ID, payment.Amount, payment.Status, payment.Method)
	if err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) GetPaymentByID(id int64) (*models.Payment, error) {
	query := `SELECT id, amount, status, method FROM payments WHERE id = $1`
	payment := &models.Payment{}
	err := r.db.QueryRow(query, id).Scan(&payment.ID, &payment.Amount, &payment.Method, &payment.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return payment, nil
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) (*models.Payment, error) {
	query := "UPDATE payments SET amount = $1, status = $2, method = $3 WHERE id = $4"
	_, err := r.db.Exec(query, payment.Amount, payment.Status, payment.Method, payment.ID)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
