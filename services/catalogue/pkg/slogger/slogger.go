package slogger

import (
	"context"
	"log/slog"
	"os"
)

// NewJsonLogger creates a new slog.Logger instance which logs to the stdout. The given level must be: "debug", "info",
// "warn" or "error". It will default to Info level if input is invalid.
func NewJsonLogger(level string) *slog.Logger {
	h := &ContextHandler{slog.NewJSONHandler(os.Stdout,
		&slog.HandlerOptions{
			Level: getLevel(level),
		},
	)}
	return slog.New(h)
}

func getLevel(lvl string) slog.Level {
	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

const slogFields = "slog_fields"

// ContextHandler embeds slog.Handler, overriding Handle method to log context attributes.
type ContextHandler struct {
	slog.Handler
}

// Handle adds contextual attributes to the Record before calling the underlying handler.
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if attrs, ok := ctx.Value(slogFields).([]slog.Attr); ok {
		for _, v := range attrs {
			r.AddAttrs(v)
		}
	}
	return h.Handler.Handle(ctx, r)
}

// AppendCtx adds a slog attribute to the provided context so that it will be included in any Record created with such
// context.
func AppendCtx(parent context.Context, attr ...slog.Attr) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	if v, ok := parent.Value(slogFields).([]slog.Attr); ok {
		v = append(v, attr...)
		return context.WithValue(parent, slogFields, v)
	}

	var v []slog.Attr
	v = append(v, attr...)
	return context.WithValue(parent, slogFields, v)
}
