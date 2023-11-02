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

	publish, err := ch.QueueDeclare(
		cfg.Broker.Publisher.Name,
		cfg.Broker.Publisher.Durable,
		cfg.Broker.Publisher.AutoDelete,
		cfg.Broker.Publisher.Exclusive,
		cfg.Broker.Publisher.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	consumer, err := ch.QueueDeclare(
		cfg.Broker.Consumer.Name,
		cfg.Broker.Consumer.Durable,
		cfg.Broker.Consumer.AutoDelete,
		cfg.Broker.Consumer.Exclusive,
		cfg.Broker.Consumer.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &config.Broker{
		Connection: conn,
		Publisher:  &config.Rabbit{Channel: ch, Queue: publish},
		Consumer:   &config.Rabbit{Channel: ch, Queue: consumer},
	}, nil
}
