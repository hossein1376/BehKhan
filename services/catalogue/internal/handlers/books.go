package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

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

func (h *Handler) getAllBooks(c *gin.Context) {
	books := h.repository.Book.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func (h *Handler) getBookByID(c *gin.Context) {
	id := c.Param("id")
	println(id)
}
