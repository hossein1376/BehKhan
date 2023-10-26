package Http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func ServeHttp(app *config.Application) {
	f := fiber.New(fiber.Config{DisableStartupMessage: true})

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Logger.Info(fmt.Sprintf("starting server on port %s", app.Settings.Http.Port))
	err := f.Listen(fmt.Sprintf(":%v", app.Settings.Http.Port))
	if err != nil {
		app.Logger.Error("failed to start HTTP server", "error", err)
		return
	}
}
