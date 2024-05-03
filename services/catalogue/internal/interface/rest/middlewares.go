package rest

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hossein1376/BehKhan/catalogue/pkg/slogger"
	"github.com/oklog/ulid/v2"
)

type middlewares struct {
	logger *slog.Logger
}

func newMiddlewares(logger *slog.Logger) *middlewares {
	return &middlewares{logger: logger}
}

func (m *middlewares) requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ulid.New(ulid.Now(), ulid.DefaultEntropy())
		if err != nil {
			m.logger.ErrorContext(c.Request.Context(), "error generating ulid", slog.Any("err", err))
			return
		}
		reqID := id.String()

		ctx := context.WithValue(c.Request.Context(), "request_id", reqID)
		ctx = slogger.WithAttrs(ctx, slog.String("request_id", reqID))

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (m *middlewares) recoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if msg := recover(); msg != nil {
				m.logger.ErrorContext(c.Request.Context(), "panic in HTTP server", slog.Any("msg", msg))
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": http.StatusText(http.StatusInternalServerError)})
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
