package services

import (
	"fmt"

	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/repositories"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type PaymentService interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByID(id int64) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) (*models.Payment, error)
}

type paymentService struct {
	paymentRepo repositories.PaymentRepository
}

func NewPaymentService(paymentRepo repositories.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *paymentService) CreatePayment(payment *models.Payment) error {
	err := s.paymentRepo.CreatePayment(payment)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}
	return nil
}

func (s *paymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	payment, err := s.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *paymentService) UpdatePayment(payment *models.Payment) (*models.Payment, error) {
	updatedPayment, err := s.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}
	return updatedPayment, nil
}
