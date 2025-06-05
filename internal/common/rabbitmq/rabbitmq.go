package rabbitmq

import (
	"transaction-worker/internal/common/logger"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQClient *RabbitMQ

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMQ(amqpURL string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (r *RabbitMQ) Close() {
	if err := r.Channel.Close(); err != nil {
		logger.Log.Error("Error closing chanel", err)
	}
	if err := r.Conn.Close(); err != nil {
		logger.Log.Error("Error closing connection:", err)
	}
}

func (r *RabbitMQ) Publish(queue string, body []byte) error {
	return r.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
