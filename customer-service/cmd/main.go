package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/nurmeden/PaymentGateway/customer-service/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err.Error())
	}

	e := echo.New()

}
