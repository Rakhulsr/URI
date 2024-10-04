package app

import (
	"github.com/Rakhulsr/go-url-shortener/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewRouter(apiHandler handler.URLHandler) *fiber.App {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", apiHandler.ShowForm)

	app.Post("/", apiHandler.SaveUrlHandler)
	app.Get("/:shortUrl", apiHandler.RetrieveUrlHandler)

	return app
}
