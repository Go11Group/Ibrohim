package redis

import (
	"product-service/config"

	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Load().RedisAddress,
		Password: config.Load().RedisPassword,
		DB:       config.Load().RedisDB,
	})
}
