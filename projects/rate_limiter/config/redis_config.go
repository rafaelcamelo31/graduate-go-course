package config

import (
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"
	REDIS_DB   = "REDIS_DB"
)

func NewRedisClient() *redis.Client {
	db, err := strconv.Atoi(os.Getenv(REDIS_DB))
	if err != nil {
		log.Fatalln("redis client init error.", err)
	}
	address := os.Getenv(REDIS_HOST) + ":" + os.Getenv(REDIS_PORT)

	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})

	return rdb
}
