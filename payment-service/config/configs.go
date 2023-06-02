package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Port     string
}
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Port string
	Host string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load("./payment-service/.env")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return Config{}, err
	}

	viper.BindEnv("Database.Host", "DATABASE_HOST")
	viper.BindEnv("Database.Port", "DATABASE_PORT")
	viper.BindEnv("Database.User", "DATABASE_USER")
	viper.BindEnv("Database.Password", "DATABASE_PASSWORD")
	viper.BindEnv("Database.Name", "DATABASE_NAME")

	viper.BindEnv("Redis.Host", "REDIS_HOST")
	viper.BindEnv("Redis.Port", "REDIS_PORT")

	viper.BindEnv("Config.Port", "PORT_APP")

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
