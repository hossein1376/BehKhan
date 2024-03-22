package command

import (
	"flag"
	"fmt"

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
	srvc := service.New(db)

	// mount routes
	// ...
	_ = srvc

	err = rest.NewServer(c.Rest.Addr).Start()
	if err != nil {
		return fmt.Errorf("start rest server: %w", err)
	}

	return nil
}
