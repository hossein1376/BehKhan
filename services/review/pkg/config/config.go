package config

import (
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/internal/repository"
)

type Application struct {
	Broker     *Broker
	Settings   *Settings
	Logger     *slog.Logger
	Repository *repository.Repository
}

type Settings struct {
	Http   http  `json:"http"`
	Grpc   grpc  `json:"grpc"`
	DB     db    `json:"db"`
	Broker queue `json:"broker"`
}

type Broker struct {
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

type queue struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	QueueName  string `json:"queue_name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"auto_delete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"no-wait"`
}
