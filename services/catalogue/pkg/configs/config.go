package configs

import (
	"log/slog"

	"github.com/hossein1376/BehKhan/catalogue/internal/repository"
)

type Application struct {
	Settings   *Settings
	Logger     *slog.Logger
	Repository *repository.Models
}

type Settings struct {
	HTTP http `yaml:"http"`
	Grpc grpc `yaml:"grpc"`
	DB   db   `yaml:"db"`
}

type http struct {
	Port int `yaml:"port"`
}

type grpc struct {
	Port int `yaml:"port"`
}

type db struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
