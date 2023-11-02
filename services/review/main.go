package main

import (
	"flag"
	"os"

	"github.com/hossein1376/BehKhan/review/cmd/Grpc"
	"github.com/hossein1376/BehKhan/review/cmd/Http"
	"github.com/hossein1376/BehKhan/review/internal/repository"
	"github.com/hossein1376/BehKhan/review/pkg/config"
	"github.com/hossein1376/BehKhan/review/pkg/database"
	"github.com/hossein1376/BehKhan/review/pkg/logging"
	"github.com/hossein1376/BehKhan/review/pkg/queue"
)

func main() {
	var cfg string
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Debug level logs")
	flag.StringVar(&cfg, "c", "configs/configs.json", "Config file path")
	flag.Parse()

	logger := logging.NewLogger(os.Stdout, debug)
	logger.Debug("logger initialized")

	settings, err := config.GetConfigs(cfg)
	if err != nil {
		logger.Error("failed to read configs", "error", err)
		return
	}
	logger.Debug("configs were loaded")

	db, disconnectDB, err := database.OpenDB(settings)
	if err != nil {
		logger.Error("failed to open database connection", "error", err)
		return
	}
	logger.Debug("opened database connection")

	broker, err := queue.NewBroker(settings)
	if err != nil {
		logger.Error("couldn't dial message queue server", "error", err)
		return
	}
	logger.Debug("successfully dialed message queue server")

	defer func() {
		if err = disconnectDB(); err != nil {
			logger.Error("failed to close database connection", "error", err)
			return
		}
		logger.Debug("closed database connection")

		if err = broker.Channel.Close(); err != nil {
			logger.Error("failed to close message queue channel", "error", err)
			return
		}
		logger.Debug("closed message queue channel")

		if err = broker.Connection.Close(); err != nil {
			logger.Error("failed to close message queue connection", "error", err)
			return
		}
		logger.Debug("closed message queue connection")
	}()

	app := &config.Application{
		Broker:     broker,
		Logger:     logger,
		Settings:   settings,
		Repository: repository.NewRepository(db),
	}

	go Grpc.ServeGrpc(app)
	Http.ServeHttp(app)
}
