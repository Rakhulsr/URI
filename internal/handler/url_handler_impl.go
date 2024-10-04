package handler

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/Rakhulsr/go-url-shortener/internal/model/web"
	"github.com/Rakhulsr/go-url-shortener/internal/service"
	"github.com/gofiber/fiber/v2"
)

type URLHandlerImpl struct {
	storageService service.StorageService
}

func NewURLHandlerImpl(storageService service.StorageService) *URLHandlerImpl {
	return &URLHandlerImpl{
		storageService: storageService,
	}
}

func (h *URLHandlerImpl) ShowForm(c *fiber.Ctx) error {
	shortUrl := c.Query("shortUrl", "")
	responseData := web.ResponseData{
		ShortUrl: shortUrl,
	}

	return c.Render("form", responseData)
}

func (h *URLHandlerImpl) SaveUrlHandler(c *fiber.Ctx) error {
	var reqBody web.ReqBody

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:    fiber.ErrBadRequest.Code,
			Status:  "Failed",
			Message: "Failed to create request",
			Data:    nil,
		})
	}

	shortUrl := h.storageService.SaveUrlMap(reqBody.OriginalUrl, reqBody.UserId)

	baseURL := os.Getenv("RAILWAY_STATIC_URL")
	if baseURL == "" {
		baseURL = os.Getenv("BASE_URL")
	}

	baseURL = strings.Trim(baseURL, "\"")
	resURL := baseURL + "/" + shortUrl

	responseData := web.ResponseData{
		OriginalUrl: reqBody.OriginalUrl,
		ShortUrl:    resURL,
	}

	c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  "Success",
		Message: "URL successfully shortened",
		Data:    responseData,
	})

	return c.Redirect("/?shortUrl="+resURL, fiber.StatusSeeOther)
}

func (h *URLHandlerImpl) RetrieveUrlHandler(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	originalUrl := h.storageService.RetrieveRealUrl(shortUrl)

	if originalUrl == "" {
		return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
			Code:    fiber.StatusNotFound,
			Status:  "Failed",
			Message: "Short URL not found",
			Data:    nil,
		})
	}

	// log.Println("Redirecting to ==> ", originalUrl)
	u, err := url.Parse(originalUrl)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:    fiber.StatusInternalServerError,
			Status:  "Error",
			Message: "Failed to redirect",
			Data:    nil,
		})
	}

	return c.Redirect(u.String(), fiber.StatusSeeOther)
}
