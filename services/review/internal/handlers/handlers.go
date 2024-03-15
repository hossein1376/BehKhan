package handlers

import (
	"log/slog"

	"github.com/hossein1376/BehKhan/review/internal/repository"
	"github.com/hossein1376/BehKhan/review/pkg/config"
	"github.com/hossein1376/BehKhan/review/pkg/transfer"
)

type handler struct {
	*slog.Logger
	*config.Rabbit
	*config.Settings
	*transfer.Response
	*repository.Repository
}

func NewHandler(app *config.Application) *handler {
	return &handler{
		Logger:     app.Logger,
		Rabbit:     app.Rabbit,
		Settings:   app.Settings,
		Repository: app.Repository,
		Response:   transfer.NewResponse(),
	}
}
