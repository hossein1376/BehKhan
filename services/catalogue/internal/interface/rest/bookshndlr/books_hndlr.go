package bookshndlr

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest/serde"
)

type BooksHndlr struct {
	Services service.Service
	Logger   *slog.Logger
}

func New(g *gin.RouterGroup, srvc service.Service, logger *slog.Logger) BooksHndlr {
	bookHandlers := BooksHndlr{
		Services: srvc,
		Logger:   logger,
	}
	g.Handle(http.MethodPost, "", bookHandlers.CreateNewBookHandler)
	g.Handle(http.MethodGet, ":id", bookHandlers.GetBookByIDHandler)

	return bookHandlers
}

func (h BooksHndlr) CreateNewBookHandler(c *gin.Context) {
	req, err := bindCreateNewBook(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = h.Services.BookSrvc.Create(c.Request.Context(), req.Title)
	if err != nil {
		h.Logger.DebugContext(c.Request.Context(), "CreateNewBook", "error", err)
		c.JSON(serde.Status(err))
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

func (h BooksHndlr) GetBookByIDHandler(c *gin.Context) {
	req, err := bindGetBookByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	resp, err := h.Services.BookSrvc.GetByID(c.Request.Context(), req.ID)
	if err != nil {
		h.Logger.DebugContext(c.Request.Context(), "GetBookByID", "error", err)
		c.JSON(serde.Status(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": resp})
	return
}
