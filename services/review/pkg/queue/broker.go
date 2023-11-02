package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func NewBroker(cfg *config.Settings) (*config.Broker, error) {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.Broker.Username,
		cfg.Broker.Password,
		cfg.Broker.Host,
		cfg.Broker.Port,
	)
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		cfg.Broker.QueueName,
		cfg.Broker.Durable,
		cfg.Broker.AutoDelete,
		cfg.Broker.Exclusive,
		cfg.Broker.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &config.Broker{Channel: ch, Queue: queue}, nil
}
