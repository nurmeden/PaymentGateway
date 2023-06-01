package services

import "github.com/nurmeden/PaymentGateway/payment-service/api/payment/repositories"

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
