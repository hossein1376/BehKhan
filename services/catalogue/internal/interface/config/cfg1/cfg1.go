package cfg1

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type db struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure,omitempty"`
}

type rest struct {
	Port string `yaml:"port"`
}

type Config struct {
	DB   db   `yaml:"db"`
	Rest rest `yaml:"rest"`
}

func (c Config) Validate() error {
	if c.DB.Name == "" {
		return fmt.Errorf("empty db name")
	}
	if c.DB.Host == "" {
		return fmt.Errorf("empty db host")
	}
	if c.DB.Port == "" {
		return fmt.Errorf("empty db port")
	}
	if c.DB.Username == "" {
		return fmt.Errorf("empty db username")
	}
	if c.Rest.Port == "" {
		return fmt.Errorf("empty rest port")
	}

	if _, err := strconv.Atoi(c.DB.Port); err != nil {
		return fmt.Errorf("invalid db port: %s", c.DB.Port)
	}
	if _, err := strconv.Atoi(c.Rest.Port); err != nil {
		return fmt.Errorf("invalid rest port: %s", c.Rest.Port)
	}

	return nil
}

func Parse(data []byte) (*Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("parsing yaml file: %w", err)
	}
	return c, nil
}
