package internal

import (
	"net/http"

	"github.com/gofiber/fiber"

	_teamHandler "soccer-api/internal/team/delivery"
)

// Setup will start the app
func Setup() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(http.StatusOK).Send("ok")
	})

	_teamHandler.New(app, nil)
	return app
}
