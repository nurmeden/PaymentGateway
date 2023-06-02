package services

import (
	"fmt"

	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/repositories"
	"github.com/nurmeden/PaymentGateway/shared/models"
)

type PaymentService interface {
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
