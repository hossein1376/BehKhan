package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/hossein1376/BehKhan/review/internal/dto"
	"github.com/hossein1376/BehKhan/review/pkg/transfer"
)

func (h *handler) createNewReviewHandler(c *fiber.Ctx) error {
	req := &dto.CreateReviewRequest{}
	err := c.BodyParser(req)
	if err != nil {
		h.Info(createNewReview, "status", transfer.BadRequest, "error", err)
		return h.BadRequestResponse(c, err)
	}

	response, err := h.Reviews.Create(req)
	if err != nil {
		h.Error(createNewReview, "status", transfer.InternalServerError, "error", err)
		return h.InternalServerErrorResponse(c, err)
	}

	h.Info(createNewReview, "status", transfer.Created, "response", response)
	return h.CreatedResponse(c, response)
}

func (h *handler) getReviewByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		h.Info(getReviewByID, "status", transfer.BadRequest, "error", "empty id parameter")
	}
	response, err := h.Reviews.Get(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.BadRequestError{}):
			h.Info(getReviewByID, "status", transfer.BadRequest, "id", id, "error", err)
			return h.BadRequestResponse(c, nil)

		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(getReviewByID, "status", transfer.NotFound, "id", id, "error", err)
			return h.NotFoundResponse(c, nil)

		default:
			h.Error(getReviewByID, "status", transfer.InternalServerError, "id", id, "error", err)
			return h.InternalServerErrorResponse(c, nil)
		}
	}

	h.Info(getReviewByID, "status", transfer.OK, "response", response)
	return h.OKResponse(c, response)
}
