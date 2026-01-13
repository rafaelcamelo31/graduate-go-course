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
	redisDb, err := strconv.Atoi(os.Getenv(REDIS_DB))
	if err != nil {
		log.Fatalln("redis client init error.", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv(REDIS_HOST) + ":" + os.Getenv(REDIS_PORT),
		DB:   redisDb,
	})

	return rdb
}
