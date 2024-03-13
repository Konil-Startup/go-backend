package redis

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type Redis struct{}

func NewCache() (*cache.Cache, error) {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Password (if used)
		DB:       0,                // Database number
	})
	// Check connection to Redis
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	mycache := cache.New(&cache.Options{
		Redis: client,
		// LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return mycache, nil
}
