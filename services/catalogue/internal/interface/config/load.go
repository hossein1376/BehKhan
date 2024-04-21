package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Load reads the config file at the given path, parse its content and validate the configurations.
func Load(path string) (Config, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading file: %w", err)
	}

	c, err := Parse(f)
	if err != nil {
		return Config{}, fmt.Errorf("parsing configs: %w", err)
	}

	if err = c.Validate(); err != nil {
		return Config{}, fmt.Errorf("validating configs: %w", err)
	}

	return c, nil
}

// Parse reads the yaml data into a Config struct. It does not perform any validations on the configurations themselves.
func Parse(data []byte) (Config, error) {
	c := Config{}
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return Config{}, fmt.Errorf("parsing yaml file: %w", err)
	}
	return c, nil
}
