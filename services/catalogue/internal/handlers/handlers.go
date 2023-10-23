package handlers

import (
	"log/slog"

	"github.com/hossein1376/BehKhan/catalogue/internal/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/configs"
	"github.com/hossein1376/BehKhan/catalogue/pkg/transfer"
)

type handler struct {
	*slog.Logger
	*transfer.Response
	*repository.Repository
}

func NewHandler(app *configs.Application) *handler {
	return &handler{
		Logger:     app.Logger,
		Repository: app.Repository,
		Response:   transfer.NewResponse(app.Logger),
	}
}
