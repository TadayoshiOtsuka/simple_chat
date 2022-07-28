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

func (s *PubSubService) Publish(ctx context.Context, channel string, payload any) {
	s.client.Publish(ctx, channel, payload)
}

func (s *PubSubService) Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	return s.client.Subscribe(ctx, channel).Channel()
}
