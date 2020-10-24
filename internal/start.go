package internal

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	_teamHandler "soccer-api/internal/team/delivery"
)

// Setup will start the app
func Setup() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("ok")
	})

	_teamHandler.New(app, nil)
	return app
}
