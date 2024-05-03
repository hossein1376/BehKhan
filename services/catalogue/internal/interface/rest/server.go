package rest

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/config"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest/bookshndlr"
)

type Server struct {
	srv *http.Server
}

func NewServer(srvc services.Service, logger *slog.Logger, cfg config.Rest) *Server {
	// create a new Engine instance
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// manage middlewares
	var middleware []gin.HandlerFunc
	m := newMiddlewares(logger)

	// add request_id if enabled
	if cfg.RequestID {
		middleware = append(middleware, m.requestIDMiddleware())
	}

	// optional logger middleware
	if cfg.Logger {
		middleware = append(middleware, m.loggerMiddleware())
	}

	// default recover middleware
	middleware = append(middleware, m.recoverMiddleware())

	// applying middlewares and create a new server
	engine.Use(middleware...)
	srv := &http.Server{
		Handler: engine,
		Addr:    cfg.Addr,
	}

	// mounting routes
	bookshndlr.New(engine.Group("books"), srvc, logger)

	return &Server{
		srv: srv,
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
