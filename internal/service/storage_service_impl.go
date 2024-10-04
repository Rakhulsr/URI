package service

import (
	"context"
	"fmt"

	"github.com/Rakhulsr/go-url-shortener/internal/db"
	"github.com/Rakhulsr/go-url-shortener/internal/helper"
	"github.com/Rakhulsr/go-url-shortener/internal/shortener"
	"github.com/google/uuid"

	"github.com/redis/go-redis/v9"
)

type StorageServiceImpl struct {
	RedisClient   *db.RedisClient
	LinkGenarator shortener.LinkGenarator
}

func NewStorageServiceImpl(redisClient *db.RedisClient, linkGenrator shortener.LinkGenarator) *StorageServiceImpl {
	return &StorageServiceImpl{
		RedisClient:   redisClient,
		LinkGenarator: linkGenrator,
	}
}

// func (s *StorageServiceImpl) SaveUrlMap(originalUrl, UserId string) string {
// 	UserId = uuid.NewString()
// 	ctx := context.Background()

// 	shortUrl := s.LinkGenarator.GenerateShortLink(originalUrl, UserId)
// 	duration := 5 * time.Hour

// 	err := s.RedisClient.Client.Set(ctx, shortUrl, originalUrl, duration).Err()

// 	helper.PanicIfError("Failed to Save URL mapping: %v", err)

// 	// fmt.Printf("Successfully Saved: \nshortURL:%s\norinalURL:%s", shortUrl, originalUrl)

// 	return shortUrl

// }

func (s *StorageServiceImpl) SaveUrlMap(originalUrl, userId string) string {
	ctx := context.Background()

	if userId == "" {
		userId = uuid.New().String()
	}

	existingShortUrl, err := s.RedisClient.Client.Get(ctx, originalUrl).Result()
	if err != nil && err != redis.Nil {
		helper.PanicIfError("Failed to retrieve existing short URL: %v", err)
	}

	if existingShortUrl != "" {
		err := s.RedisClient.Client.Del(ctx, existingShortUrl).Err()
		if err != nil {
			helper.PanicIfError("Failed to delete existing short URL: %v", err)
		}
	}
	shortUrl := s.LinkGenarator.GenerateShortLink(originalUrl, userId)
	err = s.RedisClient.Client.Set(ctx, originalUrl, shortUrl, 0).Err()
	if err != nil {
		helper.PanicIfError("Failed to create new short URL: %v", err)
	}

	err = s.RedisClient.Client.Set(ctx, shortUrl, originalUrl, 0).Err()
	if err != nil {
		helper.PanicIfError("Failed to create reverse mapping from short URL to original URL: %v", err)
	}

	return shortUrl
}

func (s *StorageServiceImpl) RetrieveRealUrl(shortUrl string) string {
	ctx := context.Background()
	result, err := s.RedisClient.Client.Get(ctx, shortUrl).Result()

	if err != nil {
		if err == redis.Nil {
			fmt.Printf("No mapping found for shortUrl: %s\n", shortUrl)
			return ""
		}
		helper.PanicIfError("Failed to retrieve URL: %v", err)
	}

	return result
}
