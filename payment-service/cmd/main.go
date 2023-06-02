package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/handlers"
	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/repositories"
	"github.com/nurmeden/PaymentGateway/payment-service/api/payment/services"
	"github.com/nurmeden/PaymentGateway/payment-service/config"
	"github.com/nurmeden/PaymentGateway/payment-service/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	db, err := database.InitDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// redisClient := cache.InitRedisClient(cfg.Redis)

	e := echo.New()

	paymentRepository := repositories.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepository)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	e.POST("/payments", paymentHandler.CreatePayment)
	// e.GET("/payments/:id", paymentHandler.GetPaymentByID)
	// e.PUT("/payments/:id", paymentHandler.UpdatePayment)
	// e.DELETE("/payments/:id", paymentHandler.DeletePayment)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
