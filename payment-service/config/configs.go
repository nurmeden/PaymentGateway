package config

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
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

}
