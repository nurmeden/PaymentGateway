package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/nurmeden/PaymentGateway/customer-service/config"
)

func InitDatabase(dbConfig config.DatabaseConfig) (*sql.DB, error) {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
