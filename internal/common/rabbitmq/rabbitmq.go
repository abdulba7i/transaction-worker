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

func (r *RabbitMQ) ConsumeMessages(queue string) (<-chan amqp091.Delivery, error) {
	msgs, err := r.Channel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQ) DeclareQueue(name string) error {
	_, err := r.Channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}
