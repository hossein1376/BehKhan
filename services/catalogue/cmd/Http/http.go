package Http

import (
	"fmt"
	"log"

	"github.com/hossein1376/BehKhan/catalogue/internal/handlers"
	"github.com/hossein1376/BehKhan/catalogue/pkg/configs"
)

func ServeHttp(app *configs.Application) {
	h := handlers.NewHandler(app)
	r := h.Router()

	app.Logger.Info(fmt.Sprintf("starting HTTP server on port %d", app.Settings.HTTP.Port))

	err := r.Run(fmt.Sprintf(":%d", app.Settings.HTTP.Port))
	if err != nil {
		log.Fatal(err)
	}
}
