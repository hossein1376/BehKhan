package logging

import (
	"io"
	"log/slog"
)

func NewLogger(dst io.Writer, verbose bool) *slog.Logger {
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}

	jsonHandler := slog.NewJSONHandler(dst, &slog.HandlerOptions{Level: level})
	logger := slog.New(jsonHandler)
	return logger
}
