package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
}

func NewCacheService(config RedisConfig) (*CacheService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &CacheService{client: client}, nil
}

func (c *CacheService) Close() error {
	return c.client.Close()
}

// Add more methods for cache operations here
