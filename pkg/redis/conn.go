package redis

import (
	"github.com/go-redis/redis/v9"
	"store/pkg/config"
)

func NewRedisClient(config config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_URL,
		Password: "",
		DB:       0,
	})

	return client
}
