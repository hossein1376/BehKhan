package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func (h *handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{DisableStartupMessage: true})

	f.Use(recover.New())

	f.Get("/", h.addNewReviewHandler)

	return f
}
