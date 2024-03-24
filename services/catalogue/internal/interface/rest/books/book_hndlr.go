package books

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/domain/dto"
)

type BookRestHndlr struct {
	Services service.Service
	Logger   *slog.Logger
}

func NewBookRestHndlr(g *gin.RouterGroup, srvc service.Service, logger *slog.Logger) BookRestHndlr {
	bookHandlers := BookRestHndlr{
		Services: srvc,
		Logger:   logger,
	}
	g.Handle(http.MethodPost, "", bookHandlers.CreateNewBookHandler)
	g.Handle(http.MethodGet, ":id", bookHandlers.GetBookByIDHandler)

	return bookHandlers
}

func (h BookRestHndlr) CreateNewBookHandler(c *gin.Context) {

}

func (h BookRestHndlr) GetBookByIDHandler(c *gin.Context) {
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
