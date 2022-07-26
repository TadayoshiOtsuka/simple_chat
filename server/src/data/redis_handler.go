package data

import (
	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	Client *redis.Client
}

func NewRedisHandler() *RedisHandler {
	c := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return &RedisHandler{
		Client: c,
	}
}
