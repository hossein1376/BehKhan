package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/hossein1376/BehKhan/catalogue/internal/interface/config"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/rest"
)

func Run() {
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/sample.yaml", "Configuration File")
	flag.Parse()

	c, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "load configs: %v", err)
		os.Exit(1)
	}

	err = rest.Start(c.Rest.Port)
	if err != nil {
		fmt.Fprintf(os.Stdout, "start rest server: %v", err)
		os.Exit(1)
	}
}
