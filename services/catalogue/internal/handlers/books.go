package handlers

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/transfer"
	"github.com/hossein1376/BehKhan/catalogue/proto/cataloguePB"
)

type Server struct {
	cataloguePB.UnimplementedBookServiceServer
}

func (s Server) GetBook(_ context.Context, in *cataloguePB.BookRequest) (*cataloguePB.BookResponse, error) {
	books := make([]*cataloguePB.Book, 0, len(in.GetId()))

	return &cataloguePB.BookResponse{
		Books: books,
	}, nil
}

func (h *handler) getAllBooks(c *gin.Context) {
	books, err := h.Book.GetAll()
	if err != nil {
		h.InternalServerErrorResponse(c, err)
		return
	}

	h.StatusCreatedResponse(c, gin.H{"books": books})
}

func (h *handler) getBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.NotFoundResponse(c, nil)
		return
	}

	book, err := h.Book.GetByID(id)
	if err != nil {
		switch {
		case errors.As(err, &transfer.NotFoundError{}):
			h.NotFoundResponse(c, err)
		default:
			h.InternalServerErrorResponse(c, err)
		}
		return
	}

	h.StatusOKResponse(c, gin.H{"book": book})
}
