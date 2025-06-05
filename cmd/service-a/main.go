package main

import (
	"os"
	"transaction-worker/internal/common/config"
	"transaction-worker/internal/common/logger"
	"transaction-worker/internal/common/rabbitmq"
	"transaction-worker/internal/service-a/handler"

	"github.com/gin-gonic/gin"
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

	// running server
	logger.Log.Info("Server is running on port :8080")

	r := gin.Default()

	rabbitClinet, err := rabbitmq.NewRabbitMQ(amqpServerURL)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer rabbitClinet.Close()

	r.POST("/transfer", handler.TransferHandler(rabbitClinet))
	r.Run(":8080")
}
