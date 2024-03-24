package command

import (
	"flag"
	"fmt"
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

	dsn := fmt.Sprintf("%s:%s@/%s", c.DB.Username, c.DB.Password, c.DB.Name)
	db, err := pool.New(dsn)
	if err != nil {
		return fmt.Errorf("open database connection: %w", err)
	}

	services := service.New(db)
	srv := rest.NewServer()
	defer srv.Stop()
	srv.Mount(services)

	// graceful stop
	startErr := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
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
