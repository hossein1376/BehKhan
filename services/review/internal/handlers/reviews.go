package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) addNewReviewHandler(c *fiber.Ctx) error {
	h.Reviews.Get()
	return h.StatusOKResponse(c, "done!")
}
