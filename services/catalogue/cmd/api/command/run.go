package command

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/services"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/migration"
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

	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load configs: %w", err)
	}

	logger := slogger.NewJsonLogger(cfg.Logger.Level)
	logger.Debug("initialized logger")

	dsn := fmt.Sprintf("%s:%s@/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Name)
	db, err := pool.New(dsn)
	if err != nil {
		return fmt.Errorf("open database connection: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("closing database connection", slog.Any("error", err))
		}
	}()
	logger.Debug("open database connection pool")

	// run DB migrations
	if err := migration.Migrate(db, cfg.Migration); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	srvc := services.New(db)

	httpSrv := rest.NewServer(srvc, logger, cfg.Rest)
	defer func() {
		logger.Debug("gracefully stopping HTTP server")
		err := httpSrv.Stop()
		if err != nil {
			logger.Error("failed to gracefully stop HTTP server", slog.Any("error", err))
		}
	}()

	grpcSrv := grpc.NewServer(srvc, logger, cfg.GRPC)
	defer func() {
		logger.Debug("gracefully stopping gRPC server")
		grpcSrv.Stop()
	}()

	// graceful stop
	startErr := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// start HTTP server
	go func() {
		logger.Info("starting HTTP server", slog.String("address", cfg.Rest.Addr))
		err := httpSrv.Start()
		startErr <- fmt.Errorf("HTTP server startup: %w", err)
	}()

	// start gRPC server
	go func() {
		logger.Info("starting gRPC server", slog.String("address", cfg.GRPC.Addr))
		err := grpcSrv.Start()
		startErr <- fmt.Errorf("gRPC server startup: %w", err)
	}()

	select {
	case err := <-startErr:
		return err
	case <-quit:
		logger.Info("received signal to stop server")
		return nil
	}
}
