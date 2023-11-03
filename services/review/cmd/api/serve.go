package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/hossein1376/BehKhan/review/cmd/Grpc"
	"github.com/hossein1376/BehKhan/review/cmd/Http"
	"github.com/hossein1376/BehKhan/review/internal/repository"
	"github.com/hossein1376/BehKhan/review/pkg/brokers"
	"github.com/hossein1376/BehKhan/review/pkg/config"
	"github.com/hossein1376/BehKhan/review/pkg/database"
	"github.com/hossein1376/BehKhan/review/pkg/logging"
)

func serve() {
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

	rabbit, err := brokers.OpenRabbit(settings)
	if err != nil {
		logger.Error("couldn't dial RabbitMQ server", "error", err)
		return
	}
	logger.Debug("successfully dialed RabbitMQ server")

	defer func() {
		if err = disconnectDB(); err != nil {
			logger.Error("failed to close database connection", "error", err)
			return
		}
		logger.Debug("closed database connection")

		if err = rabbit.Close(); err != nil {
			logger.Error("failed to close RabbitMQ connection", "error", err)
			return
		}
		logger.Debug("closed RabbitMQ connection")
	}()

	signals := config.Signals{ShutdownHTTP: make(chan os.Signal), ShutdownGRPC: make(chan os.Signal)}

	app := &config.Application{
		Rabbit:     rabbit,
		Logger:     logger,
		Settings:   settings,
		Repository: repository.NewRepository(db),
		Signals:    signals,
	}

	go Grpc.ServeGrpc(app)
	go Http.ServeHttp(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	interrupt := <-quit

	app.Logger.Info("shutting down server", "signal", interrupt.String())
	app.Signals.ShutdownHTTP <- interrupt
	app.Signals.ShutdownGRPC <- interrupt
	return
}
