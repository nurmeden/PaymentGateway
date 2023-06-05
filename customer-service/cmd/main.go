package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/repositories"
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/services"
	"github.com/nurmeden/PaymentGateway/customer-service/api/payment/services"
	"github.com/nurmeden/PaymentGateway/customer-service/config"
	"github.com/nurmeden/PaymentGateway/customer-service/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err.Error())
	}

	e := echo.New()

	db, err := database.InitDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to the database: %s", err.Error())
	}
	defer db.Close()

	customerRepo := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepo)
}
