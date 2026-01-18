package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"
	REDIS_DB   = "REDIS_DB"
)

func NewRedisClient() *redis.Client {
	address := os.Getenv(REDIS_HOST) + ":" + os.Getenv(REDIS_PORT)

	rdb := redis.NewClient(&redis.Options{
		Addr: address,
	})

	return rdb
}
