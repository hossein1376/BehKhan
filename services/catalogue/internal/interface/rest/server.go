package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest/books"
)

type Server struct {
	srv *gin.Engine
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	return &Server{engine}
}

func (s *Server) Start(addr string) error {
	return s.srv.Run(addr)
}

func (s *Server) Mount(srvc service.Service) {
	books.NewBookRestHndlr(s.srv.Group("books"), srvc)
}
