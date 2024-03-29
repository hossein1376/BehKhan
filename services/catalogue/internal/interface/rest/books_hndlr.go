package rest

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
)

type BooksHndlr struct {
	Services services.Service
	Logger   *slog.Logger
}

func NewBooksHndlr(g *gin.RouterGroup, srvc services.Service, logger *slog.Logger) BooksHndlr {
	bookHandlers := BooksHndlr{
		Services: srvc,
		Logger:   logger,
	}
	g.Handle(http.MethodPost, "", bookHandlers.CreateNewBookHandler)
	g.Handle(http.MethodGet, ":id", bookHandlers.GetBookByIDHandler)

	return bookHandlers
}

func (h BooksHndlr) CreateNewBookHandler(c *gin.Context) {

}

func (h BooksHndlr) GetBookByIDHandler(c *gin.Context) {
	req := &dto.GetBookByIDRequest{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.Services.BookSrvc.GetByID(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": resp})
	return
}