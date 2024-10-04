package db

import (
	"context"
	"fmt"
	"os"

	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {

	godotenv.Load()

	residURL := os.Getenv("REDIS_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     residURL,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed connect to Redis: %v", err))
	}

	return &RedisClient{Client: client}

}
