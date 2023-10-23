package handlers

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/configs"
)

type Handler struct {
	repository *repository.Models
	logger     *slog.Logger
}

func NewHandler(app *configs.Application) Handler {
	return Handler{
		repository: app.Repository,
		logger:     app.Logger,
	}
}

func (h *Handler) Router() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Recovery())

	/*	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})*/

	base := r.Group("/api/v1")

	books := base.Group("/books")
	books.GET("/", h.getAllBooks)
	books.GET("/:id", h.getBookByID)

	return r
}
