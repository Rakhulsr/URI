package service

import (
	"testing"

	"github.com/Rakhulsr/go-url-shortener/internal/db"
	"github.com/Rakhulsr/go-url-shortener/internal/shortener"
	"github.com/stretchr/testify/assert"
)

var testStorageService *StorageServiceImpl

func init() {

	redisClient := db.NewRedisClient()

	testStorageService = NewStorageServiceImpl(redisClient, shortener.NewLinkGeneratorImpl())

}

func TestInitStrorageService(t *testing.T) {
	assert.NotNil(t, testStorageService.RedisClient, "Redis client should not be nil")

}

func TestInsertURL(t *testing.T) {
	originalUrl := "https://www.samsung.com/id/smartphones/galaxy-a/galaxy-a54-5g-green-256gb-sm-a546elgdxid/"
	userUUID := "9b898b62-0391-44c0-ab59-c7eea318bc10"
	shortUrl := testStorageService.LinkGenarator.GenerateShortLink(originalUrl, userUUID)

	testStorageService.SaveUrlMap(originalUrl, userUUID)
	retrievedUrl := testStorageService.RetrieveRealUrl(shortUrl)

	assert.Equal(t, originalUrl, retrievedUrl)

}
