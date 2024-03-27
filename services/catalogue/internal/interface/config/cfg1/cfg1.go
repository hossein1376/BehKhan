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
	Addr string `yaml:"addr"`
}

type grpc struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	DB   db   `yaml:"db"`
	Rest rest `yaml:"rest"`
	Grpc grpc `yaml:"grpc"`
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
	if c.Rest.Addr == "" {
		return fmt.Errorf("empty Rest address")
	}
	if c.Grpc.Addr == "" {
		return fmt.Errorf("empty gRPC address")
	}
	if _, err := strconv.Atoi(c.DB.Port); err != nil {
		return fmt.Errorf("invalid db port: %s", c.DB.Port)
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
