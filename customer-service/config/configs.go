package config

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
