package main

import (
	"flag"
	"os"

	"github.com/hossein1376/BehKhan/catalogue/cmd/Grpc"
	"github.com/hossein1376/BehKhan/catalogue/cmd/Http"
	"github.com/hossein1376/BehKhan/catalogue/internal/repository"
	"github.com/hossein1376/BehKhan/catalogue/pkg/config"
	"github.com/hossein1376/BehKhan/catalogue/pkg/database"
	"github.com/hossein1376/BehKhan/catalogue/pkg/logging"
)

// @title           Catalogue
// @version         0.1.0
// @description     BehKhan's catalogue microservice.
// @contact.name    Hossein Yazdani
// @contact.url     https://GodlyNice.ir
// @license.name    MIT license
// @license.url     https://opensource.org/license/mit/
// @host            localhost:8002
// @BasePath        /api/v1/catalogue
func main() {
	var debug bool
	var cfg string
	flag.BoolVar(&debug, "debug", false, "Debug level logs")
	flag.StringVar(&cfg, "c", "./config/config.yaml", "Config file path")
	flag.Parse()

	logger := logging.NewLogger(os.Stdout, debug)

	settings, err := config.GetSettings(cfg)
	if err != nil {
		logger.Error("failed to read the settings", "error", err)
		return
	}

	db, err := database.GetDB(settings)
	if err != nil {
		logger.Error("failed to open database", "error", err)
		return
	}

	app := &config.Application{
		Settings:   settings,
		Logger:     logger,
		Repository: repository.NewRepository(db),
	}

	go Grpc.ServeGrpc(app)
	Http.ServeHttp(app)
}
