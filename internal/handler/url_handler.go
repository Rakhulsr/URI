package handler

import "github.com/gofiber/fiber/v2"

type URLHandler interface {
	SaveUrlHandler(c *fiber.Ctx) error
	RetrieveUrlHandler(c *fiber.Ctx) error
	ShowForm(c *fiber.Ctx) error
}
