package services

import (
	"context"
	"encoding/json"
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

type UpdateReviewMsg struct {
	Total   int64   `json:"total"`
	Average float64 `json:"average"`
}

func CatalogueReviewsUpdate(rabbit *config.Rabbit, total int64, average float64) error {
	if rabbit == nil {
		return errors.New("no connection to RabbitMQ available")
	}

	body := UpdateReviewMsg{
		Total:   total,
		Average: average,
	}
	msg, err := json.Marshal(body)
	if err != nil {
		return err
	}

	err = rabbit.Channel.PublishWithContext(context.Background(),
		"",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)

	return err
}
