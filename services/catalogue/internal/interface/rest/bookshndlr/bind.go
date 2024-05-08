package bookshndlr

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/entity"
)

type getBookByIDRequest struct {
	ID entity.BookID
}

type createNewBookRequest struct {
	Title string `json:"title"`
}

func bindGetBookByID(c *gin.Context) (*getBookByIDRequest, error) {
	idParam := c.Param("id")
	if idParam == "" {
		return nil, fmt.Errorf("id parameter must be provided")
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("id parameter must be a valid, positive number")
	}

	return &getBookByIDRequest{ID: entity.BookID(id)}, nil
}

func bindCreateNewBook(c *gin.Context) (*createNewBookRequest, error) {
	var req *createNewBookRequest
	if err := c.BindJSON(&req); err != nil {
		return nil, fmt.Errorf("biding request: %w", err)
	}

	if req.Title == "" {
		return nil, fmt.Errorf("title must be provided")
	}

	return req, nil
}
