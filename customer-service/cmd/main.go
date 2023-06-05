package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/handlers"
	"github.com/nurmeden/PaymentGateway/customer-service/api/customer/repositories"
	customerServices "github.com/nurmeden/PaymentGateway/customer-service/api/customer/services"
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
	customerService := customerServices.NewCustomerService(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerService)

	e.POST("/customers", customerHandler.CreateCustomer)
	e.GET("/customers/:id", customerHandler.GetCustomerByID)
	e.PUT("/customers", customerHandler.UpdateCustomer)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server started on %s", addr)
	e.Start(addr)
}
