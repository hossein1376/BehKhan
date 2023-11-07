package handlers

import (
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/pkg/config"
	"github.com/hossein1376/BehKhan/catalogue/pkg/transfer"
	"github.com/hossein1376/BehKhan/structure/cataloguePB"
)

type Server struct {
	*config.Application
	cataloguePB.UnimplementedBookServiceServer
}

func (s Server) GetBook(_ context.Context, req *cataloguePB.BookRequest) (*cataloguePB.BookResponse, error) {
	resp, err := s.Application.Repository.Book.GetByID(req.GetId()...)
	if err != nil {
		return nil, err
	}

	books := make([]*cataloguePB.Book, 0, len(resp))
	for _, book := range resp {
		books = append(books, &cataloguePB.Book{Id: book.ID, Name: book.Name})
	}

	return &cataloguePB.BookResponse{Books: books}, nil
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

	h.StatusOKResponse(c, gin.H{"book": book[0]})
}
