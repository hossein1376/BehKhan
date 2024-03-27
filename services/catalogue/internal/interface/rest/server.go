package rest

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
)

type Server struct {
	engine *gin.Engine
	srv    *http.Server
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	srv := &http.Server{
		Handler: engine,
	}

	return &Server{
		engine: engine,
		srv:    srv,
	}
}

func (s *Server) Start(addr string) error {
	s.srv.Addr = addr
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}

func (s *Server) Mount(srvc services.Service, logger *slog.Logger) {
	NewBooksHndlr(s.engine.Group("books"), srvc, logger)
}
