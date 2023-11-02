package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	createNewReview = "Create New Review"
	getReviewByID   = "Get Review By ID"
)

func (h *handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{DisableStartupMessage: true})

	f.Use(recover.New())

	f.Get("/:id", h.getReviewByIDHandler)
	f.Post("/", h.createNewReviewHandler)

	return f
}
