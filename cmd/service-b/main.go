package main

import (
	"os"
	"transaction-worker/internal/common/config"
	"transaction-worker/internal/common/logger"
	"transaction-worker/internal/common/rabbitmq"

	"github.com/joho/godotenv"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.MustLoad()
	logger.Log.Infof("env: %s", cfg.Env)

	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	rabbitClinet, err := rabbitmq.NewRabbitMQ(amqpServerURL)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer rabbitClinet.Close()

	if err := rabbitClinet.DeclareQueue("QueueService1"); err != nil {
		logger.Log.Fatal("Failed to declare queue: ", err)
	}

	messages, err := rabbitClinet.ConsumeMessages("QueueService1")

	if err != nil {
		logger.Log.Fatalf("Error consuming messages: %v", err)
	}

	logger.Log.Info("Successfully connected to RabbitMQ")
	logger.Log.Info("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			logger.Log.Infof(" > Received message: %s\n", message.Body)

			message.Ack(false)
		}
	}()

	<-forever
}
