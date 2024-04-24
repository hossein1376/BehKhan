// Package config is home to the application's configuration.
package config

import (
	"fmt"
	"strconv"
)

// Config stores configurations of the application. It is created on startup from the configuration file, and is meant
// to remain read-only.
type Config struct {
	DB     DB     `yaml:"db"`
	Rest   Rest   `yaml:"rest"`
	GRPC   GRPC   `yaml:"grpc"`
	Logger Logger `yaml:"logger"`
}

type DB struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure,omitempty"`
}

type Rest struct {
	Addr      string `yaml:"addr"`
	Logger    bool   `yaml:"logger"`
	RequestID bool   `yaml:"request_id"`
}

type GRPC struct {
	Addr      string `yaml:"addr"`
	Logger    bool   `yaml:"logger"`
	RequestID bool   `yaml:"request_id"`
}

type Logger struct {
	Level string `yaml:"level"`
}

// Validate checks for the integrity of the provided configs.
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
	if c.GRPC.Addr == "" {
		return fmt.Errorf("empty gRPC address")
	}
	if _, err := strconv.Atoi(c.DB.Port); err != nil {
		return fmt.Errorf("invalid db port: %s", c.DB.Port)
	}

	return nil
}
