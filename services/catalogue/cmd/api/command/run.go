package command

import (
	"flag"
	"fmt"

	"github.com/hossein1376/BehKhan/catalogue/internal/application/controller"
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
	ctrl := controller.New(db)

	// mount routes
	// ...
	_ = ctrl

	err = rest.Start(c.Rest.Addr)
	if err != nil {
		return fmt.Errorf("start rest server: %w", err)
	}

	return nil
}
