package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
	}
}

func (r *RedisService) SetValue(cxt context.Context, key string, value string) error {
	return r.client.Set(cxt, key, value, 0).Err()
}
