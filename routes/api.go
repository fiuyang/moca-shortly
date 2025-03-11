package routes

import (
	"github.com/gofiber/fiber/v2"
	"scylla/handler"
)

func NewRoutesV1(
	shortenHandler *handler.ShortenHandler,
) *fiber.App {
	app := fiber.New()
	routes := app.Group("")

	// Redirect short URL
	routes.Get("/:shortCode", shortenHandler.RedirectURL)

	// Group shorten API
	shortenRouter := routes.Group("/shorten")
	shortenRouter.Post("", shortenHandler.ShortenURL)
	shortenRouter.Get("/:shortCode", shortenHandler.ShowOriginalURL)

	return app
}
