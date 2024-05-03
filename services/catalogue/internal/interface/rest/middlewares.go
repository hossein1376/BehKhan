package rest

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/pkg/reqID"
	"github.com/hossein1376/BehKhan/catalogue/pkg/slogger"
)

type middlewares struct {
	logger *slog.Logger
}

func newMiddlewares(logger *slog.Logger) *middlewares {
	return &middlewares{logger: logger}
}

func (m *middlewares) requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := reqID.NewRequestID()
		if err != nil {
			m.logger.ErrorContext(c.Request.Context(), "error generating request id", slog.Any("error", err))
			return
		}

		// put request id inside context
		ctx := context.WithValue(c.Request.Context(), reqID.RequestIDKey, id)

		// include request_id in logs
		ctx = slogger.WithAttrs(ctx, slog.String("request_id", id))

		// replace the context with the new one
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (m *middlewares) recoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				m.logger.ErrorContext(c.Request.Context(), "panic in HTTP server", slog.Any("error", msg))
				c.AbortWithStatusJSON(http.StatusInternalServerError,
					gin.H{"message": http.StatusText(http.StatusInternalServerError)})
			}
		}()
		c.Next()
	}
}

func (m *middlewares) loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		defer func() {
			if raw != "" {
				path = path + "?" + raw
			}
			m.logger.InfoContext(c.Request.Context(), "http server",
				slog.Group("request",
					slog.String("client_ip", c.ClientIP()),
					slog.String("method", c.Request.Method),
					slog.String("request_path", path),
				),
				slog.Group("response",
					slog.Int("status", c.Writer.Status()),
					slog.String("time_took", time.Since(start).String()),
				),
			)
		}()
		c.Next()
	}
}
