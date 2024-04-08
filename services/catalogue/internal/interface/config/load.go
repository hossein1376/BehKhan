package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	c, err := Parse(f)
	if err != nil {
		return nil, fmt.Errorf("parsing configs: %w", err)
	}

	if err = c.Validate(); err != nil {
		return nil, fmt.Errorf("validating configs: %w", err)
	}

	return c, nil
}

func Parse(data []byte) (*Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("parsing yaml file: %w", err)
	}
	return c, nil
}
