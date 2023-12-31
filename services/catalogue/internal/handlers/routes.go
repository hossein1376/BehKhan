package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *handler) Router() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Recovery())

	base := r.Group("/api/v1")

	books := base.Group("/books")
	books.GET("/", h.getAllBooks)
	books.GET("/:id", h.getBookByID)

	return r
}
