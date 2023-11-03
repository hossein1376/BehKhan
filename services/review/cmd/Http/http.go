package Http

import (
	"fmt"
	"syscall"

	"github.com/hossein1376/BehKhan/review/internal/handlers"
	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func ServeHttp(app *config.Application) {
	h := handlers.NewHandler(app)
	f := h.Router()

	// graceful shutdown
	go func() {
		<-app.Signals.ShutdownHTTP
		app.Logger.Debug("HTTP graceful shutdown started")

		err := f.Shutdown()
		if err != nil {
			app.Logger.Error("HTTP graceful shutdown failed")
			return
		}

		app.Logger.Debug("HTTP graceful shutdown finished")
		app.Signals.ShutdownHTTP <- syscall.Signal(0)
	}()

	app.Logger.Info(fmt.Sprintf("starting HTTP server on port %s", app.Settings.Http.Port))
	err := f.Listen(fmt.Sprintf(":%v", app.Settings.Http.Port))
	if err != nil {
		app.Logger.Error("failed to start HTTP server", "error", err)
		return
	}
}
