package command

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/config"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/grpc"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest"
	"github.com/hossein1376/BehKhan/catalogue/pkg/slogger"
)

func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/sample.yaml", "Configuration File")
	flag.Parse()

	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load configs: %w", err)
	}

	logger := slogger.NewJsonLogger(c.Logger.Level)
	logger.Debug("initialized logger")

	dsn := fmt.Sprintf("%s:%s@/%s", c.DB.Username, c.DB.Password, c.DB.Name)
	db, err := pool.New(dsn)
	if err != nil {
		return fmt.Errorf("open database connection: %w", err)
	}
	defer func() {
		if err := db.GetDB().Close(); err != nil {
			logger.Error("closing database connection", slog.Any("error", err))
		}
	}()
	logger.Debug("open database connection pool")

	services := service.New(db)

	httpSrv := rest.NewServer()
	defer func() {
		logger.Debug("gracefully stopping HTTP server")
		err := httpSrv.Stop()
		if err != nil {
			logger.Error("failed to gracefully stop HTTP server", slog.Any("error", err))
		}
	}()
	httpSrv.Mount(services, logger)

	grpcSrv := grpc.NewServer()
	defer func() {
		logger.Debug("gracefully stopping gRPC server")
		grpcSrv.Stop()
	}()
	grpcSrv.Mount(services, logger)

	// graceful stop
	startErr := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// start HTTP server
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic in HTTP server goroutine", slog.Any("msg", err))
			}
		}()
		logger.Info("starting HTTP server", slog.String("address", c.Rest.Addr))
		err := httpSrv.Start(c.Rest.Addr)
		startErr <- fmt.Errorf("HTTP server startup: %w", err)
	}()

	// start gRPC server
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("panic in gRPC server goroutine", slog.Any("msg", err))
			}
		}()
		logger.Info("starting gRPC server", slog.String("address", c.Grpc.Addr))
		err := grpcSrv.Start(c.Grpc.Addr)
		startErr <- fmt.Errorf("gRPC server startup: %w", err)
	}()

	select {
	case err := <-startErr:
		logger.Error("failed to start server", slog.Any("error", err))
		return err
	case <-quit:
		logger.Info("received signal to stop server")
		return nil
	}
}
