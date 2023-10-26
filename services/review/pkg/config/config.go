package config

import (
	"log/slog"

	"github.com/hossein1376/BehKhan/review/internal/repository"
)

type Application struct {
	Settings   *Settings
	Logger     *slog.Logger
	Repository *repository.Repository
}

type Settings struct {
	Http http `json:"http"`
	Grpc grpc `json:"grpc"`
	DB   db   `json:"db"`
}

type http struct {
	Port string `json:"port"`
}

type grpc struct {
	Port string `json:"port"`
}

type db struct {
	Name       string `json:"name"`
	Collection string `json:"collection"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}
