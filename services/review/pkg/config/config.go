package config

import (
	"log/slog"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/internal/repository"
	"github.com/hossein1376/BehKhan/review/pkg/duration"
)

type Application struct {
	Rabbit     *Rabbit
	Settings   *Settings
	Logger     *slog.Logger
	Repository *repository.Repository
	Signals    Signals
}

type Settings struct {
	Http     http     `json:"http"`
	Grpc     grpc     `json:"grpc"`
	DB       db       `json:"db"`
	Rabbit   rabbit   `json:"rabbitmq"`
	Services services `json:"services"`
}

type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
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

type rabbit struct {
	Username     string            `json:"username"`
	Password     string            `json:"password"`
	Host         string            `json:"host"`
	Port         string            `json:"port"`
	RetryTimeout duration.Duration `json:"retry_timeout"`

	QueueName  string `json:"queue_name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"no-wait"`
}

type Signals struct {
	ShutdownHTTP chan os.Signal
	ShutdownGRPC chan os.Signal
}
