package delivery

import (
	"soccer-api/internal/team"

	"github.com/gofiber/fiber/v2"
)

type teamHandler struct {
	svc team.Team
}

// New return team handler
func New(f *fiber.App, svc team.Team) {
	h := &teamHandler{svc: svc}

	f.Get("/teams", h.Fetch)
	f.Get("/teams/:id", h.Get)
	f.Post("/teams", h.Insert)
}

func (h teamHandler) Fetch(c *fiber.Ctx) error {
	return c.JSON("fetch")
}

func (h teamHandler) Get(c *fiber.Ctx) error {
	return c.JSON("get")
}

func (h teamHandler) Insert(c *fiber.Ctx) error {
	return c.JSON("insert")
}
