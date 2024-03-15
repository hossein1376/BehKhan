package brokers

import (
	"fmt"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func OpenRabbit(cfg *config.Settings, logger *slog.Logger) (*config.Rabbit, error) {
	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.Rabbit.Username,
		cfg.Rabbit.Password,
		cfg.Rabbit.Host,
		cfg.Rabbit.Port,
	)

	attempt := func() (*config.Rabbit, error) {
		conn, err := amqp.Dial(dsn)
		if err != nil {
			return nil, err
		}

		ch, err := conn.Channel()
		if err != nil {
			return nil, err
		}

		// TODO: create one for each service
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

	var (
		resp    *config.Rabbit
		err     error
		timeout = cfg.Rabbit.RetryTimeout
	)
	for {
		resp, err = attempt()
		if err == nil {
			break
		}
		logger.Error(fmt.Sprintf("Failed to connect to RabbitMQ, sleeping for %s...", timeout.String()), "error", err)
		time.Sleep(timeout.Duration)
	}

	return resp, nil
}
