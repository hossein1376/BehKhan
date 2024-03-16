package config

import (
	"fmt"
	"os"

	"github.com/hossein1376/BehKhan/catalogue/internal/interface/config/cfg1"
)

func Load(path string) (*cfg1.Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	c, err := cfg1.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("parsing configs: %w", err)
	}

	if err = c.Validate(); err != nil {
		return nil, fmt.Errorf("validating configs: %w", err)
	}

	return c, nil
}
