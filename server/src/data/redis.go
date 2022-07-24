package data

import "github.com/go-redis/redis/v8"

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	c := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		Client: c,
	}
}
