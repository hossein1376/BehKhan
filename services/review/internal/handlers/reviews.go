package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/hossein1376/BehKhan/review/internal/dto"
	"github.com/hossein1376/BehKhan/review/pkg/transfer"
)

func (h *handler) createNewReviewHandler(c *fiber.Ctx) error {
	req := &dto.CreateReviewRequest{}
	err := c.BodyParser(req)
	if err != nil {
		h.Info(createNewReview, "status", transfer.StatusBadRequest, "error", err)
		return h.BadRequestResponse(c, err)
	}

	response, err := h.Reviews.Create(req)
	if err != nil {
		h.Error(createNewReview, "status", transfer.StatusInternalServerError, "error", err)
		return h.InternalServerErrorResponse(c, err)
	}

	h.Info(createNewReview, "status", transfer.StatusCreated, "response", response)
	return h.CreatedResponse(c, response)
}

func (h *handler) getReviewByIDHandler(c *fiber.Ctx) error {
	rid := c.Params("review_id")
	bid, err := strconv.ParseInt(c.Params("book_id"), 10, 64)
	if rid == "" || err != nil {
		h.Info(getReviewByID, "status", transfer.StatusBadRequest, "error", "bad id parameter")
		return h.BadRequestResponse(c)
	}

	response, err := h.Reviews.Get(bid, rid)
	if err != nil {
		switch {
		case errors.As(err, &transfer.BadRequestError{}):
			h.Info(getReviewByID, "status", transfer.StatusBadRequest, "error", err, "book_id", bid, "review_id", rid)
			return h.BadRequestResponse(c, nil)

		case errors.As(err, &transfer.NotFoundError{}):
			h.Info(getReviewByID, "status", transfer.StatusNotFound, "error", err, "book_id", bid, "review_id", rid)
			return h.NotFoundResponse(c, nil)

		default:
			h.Error(getReviewByID, "status", transfer.StatusInternalServerError, "error", err, "book_id", bid, "review_id", rid)
			return h.InternalServerErrorResponse(c, nil)
		}
	}

	h.Info(getReviewByID, "status", transfer.StatusOK, "response", response)
	return h.OkResponse(c, response)
}
