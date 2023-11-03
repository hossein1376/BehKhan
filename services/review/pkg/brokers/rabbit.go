package brokers

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func OpenRabbit(cfg *config.Settings) (*config.Rabbit, error) {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.Rabbit.Username,
		cfg.Rabbit.Password,
		cfg.Rabbit.Host,
		cfg.Rabbit.Port,
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
		cfg.Rabbit.QueueName,
		cfg.Rabbit.Durable,
		cfg.Rabbit.AutoDelete,
		cfg.Rabbit.Exclusive,
		cfg.Rabbit.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &config.Rabbit{
		Connection: conn,
		Channel:    ch,
		Queue:      publish,
	}, nil
}
