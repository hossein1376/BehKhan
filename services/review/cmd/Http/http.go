package Http

import (
	"fmt"

	"github.com/hossein1376/BehKhan/review/internal/handlers"
	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func ServeHttp(app *config.Application) {
	h := handlers.NewHandler(app)
	f := h.Router()

	go func() {
		<-app.Signals.ShutdownHTTP
		err := f.Shutdown()
		if err != nil {
			app.Logger.Error("HTTP graceful shutdown failed")
			return
		}
	}()

	app.Logger.Info(fmt.Sprintf("starting server on port %s", app.Settings.Http.Port))
	err := f.Listen(fmt.Sprintf(":%v", app.Settings.Http.Port))
	if err != nil {
		app.Logger.Error("failed to start HTTP server", "error", err)
		return
	}
}
