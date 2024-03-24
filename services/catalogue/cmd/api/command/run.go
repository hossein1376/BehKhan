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
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest"
)

func Run() error {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/sample.yaml", "Configuration File")
	flag.Parse()

	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load configs: %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Debug("initialized logger")

	dsn := fmt.Sprintf("%s:%s@/%s", c.DB.Username, c.DB.Password, c.DB.Name)
	db, err := pool.New(dsn)
	if err != nil {
		return fmt.Errorf("open database connection: %w", err)
	}
	logger.Debug("open database connection pool")

	services := service.New(db)
	srv := rest.NewServer()
	defer func() {
		err := srv.Stop()
		if err != nil {
			logger.Error("failed to gracefully stop HTTP server", "error", err)
		}
	}()

	srv.Mount(services, logger)

	// graceful stop
	startErr := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("starting HTTP server", "address", c.Rest.Addr)
		err = srv.Start(c.Rest.Addr)
		startErr <- fmt.Errorf("HTTP server startup: %w", err)
	}()

	select {
	case err = <-startErr:
		return err
	case <-quit:
		return nil
	}
}
