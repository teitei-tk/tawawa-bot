package config

import (
	"os"

	"gopkg.in/redis.v5"
)

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return client
}
