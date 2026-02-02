package redis

import (
	"monorepo/apps/gas/app/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password, // no password set
		DB:       config.Redis.DB,       // use default DB
	})
	return rdb
}
