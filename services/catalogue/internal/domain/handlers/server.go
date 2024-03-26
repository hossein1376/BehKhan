package handlers

import (
	"log/slog"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
)

type Server interface {
	Start(string) error
	Stop() error
	Mount(services.Service, *slog.Logger)
}
