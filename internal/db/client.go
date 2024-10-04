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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed connect to Redis: %v", err))
	}

	return &RedisClient{Client: client}

}
