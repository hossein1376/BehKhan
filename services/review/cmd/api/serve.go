package main

import (
	"context"
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
	debug, cfg := false, "configs/configs.json"
	flag.BoolVar(&debug, "debug", debug, "Debug level logs")
	flag.StringVar(&cfg, "c", cfg, "Config file path")
	flag.Parse()

	logger := logging.NewLogger(os.Stdout, debug)
	logger.Debug("logger initialized")

	settings, err := config.GetConfigs(cfg)
	if err != nil {
		logger.Error("failed to read configs", "error", err)
		return
	}
	logger.Debug("configs were loaded")

	client, db, err := database.OpenDB(settings)
	if err != nil {
		logger.Error("failed to open database connection", "error", err)
		return
	}
	logger.Debug("opened database connection")

	app := &config.Application{
		Logger:     logger,
		Settings:   settings,
		Repository: repository.NewRepository(client, db),
	}

	go func() {
		app.Rabbit, err = brokers.OpenRabbit(settings, logger)
		if err != nil {
			logger.Error("couldn't dial RabbitMQ server", "error", err)
			return
		}
		logger.Debug("successfully dialed RabbitMQ server")
	}()

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			logger.Error("failed to close database connection", "error", err)
			return
		}
		logger.Debug("closed database connection")

		if app.Rabbit != nil {
			if err = app.Rabbit.Close(); err != nil {
				logger.Error("failed to close RabbitMQ connection", "error", err)
				return
			}
			logger.Debug("closed RabbitMQ connection")
		}
	}()

	app.Signals = config.Signals{ShutdownHTTP: make(chan os.Signal), ShutdownGRPC: make(chan os.Signal)}

	go Grpc.ServeGrpc(app)
	go Http.ServeHttp(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// wait indefinitely until exit signal is sent
	interrupt := <-quit
	logger.Info("shutting down the server", "signal", interrupt.String())

	// sending interrupt signal to goroutines
	app.Signals.ShutdownGRPC <- interrupt
	app.Signals.ShutdownHTTP <- interrupt

	// waiting for each one's response
	<-app.Signals.ShutdownGRPC
	<-app.Signals.ShutdownHTTP
	return
}
