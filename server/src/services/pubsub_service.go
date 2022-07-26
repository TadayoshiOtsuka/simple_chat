package services

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type PubSubService struct {
	client *redis.Client
}

func NewPubSubService() *PubSubService {
	c := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &PubSubService{
		client: c,
	}
}

func (h *PubSubService) Publish(ctx context.Context, channel string, payload any) {
	h.client.Publish(ctx, channel, payload)
}

func (h *PubSubService) Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	return h.client.Subscribe(ctx, channel).Channel()
}
