package cache

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/nurmeden/PaymentGateway/payment-service/config"
)

func initRedisClient(redisConfig config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return client
}
