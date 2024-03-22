package rest

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	addr   string
}

func NewServer(addr string) *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	return &Server{
		engine: engine,
		addr:   addr,
	}
}

func (s *Server) Start() error {
	return s.engine.Run()
}
