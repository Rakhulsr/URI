package db

import (
	"context"
	"fmt"

	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed connect to Redis: %v", err))
	}

	fmt.Println("Successfully connected to Redis")
	return &RedisClient{Client: client}

}
