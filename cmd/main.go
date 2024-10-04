package main

import (
	"log"

	"github.com/Rakhulsr/go-url-shortener/cmd/app"
	"github.com/Rakhulsr/go-url-shortener/internal/db"
	"github.com/Rakhulsr/go-url-shortener/internal/handler"
	"github.com/Rakhulsr/go-url-shortener/internal/service"
	"github.com/Rakhulsr/go-url-shortener/internal/shortener"
)

func main() {
	redisClient := db.NewRedisClient()
	linkGen := shortener.NewLinkGeneratorImpl()
	storageServiceImpl := service.NewStorageServiceImpl(redisClient, linkGen)
	storageHandlerImpl := handler.NewURLHandlerImpl(storageServiceImpl)

	router := app.NewRouter(storageHandlerImpl)

	log.Fatal(router.Listen(":8080"))
}
