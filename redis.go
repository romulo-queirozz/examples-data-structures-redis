package redis

import (
	"github.com/go-redis/redis/v9"
)

func MyRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
