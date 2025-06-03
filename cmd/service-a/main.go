package main

import (
	"net/http"
	"os"
	"transaction-worker/internal/common/config"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	cfg := config.MustLoad()
	logger.Info(cfg.Env)

	logger.Info("starting transaction-worker",
		zap.String("env", cfg.Env),
		zap.String("version", version),
	)

	// connect rabbitmq
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		logger.Panic("Error connect", zap.Error(err))
	}

	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		logger.Panic("Error create chan", zap.Error(err))
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)

	if err != nil {
		logger.Panic("Error QueueDeclare", zap.Error(err))
	}

	// running server
	logger.Info("Server is running on port :8080")

	r := gin.New()
	r.Use(
		gin.Logger(),
	)

	r.POST("/input", func(ctx *gin.Context) {
		// Структура для приёма JSON
		var json struct {
			Msg string `json:"msg"`
		}

		// Парсим JSON из тела запроса
		if err := ctx.BindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		// Проверяем, что поле msg не пустое
		if json.Msg == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "msg field is required"})
			return
		}

		// Создаём сообщение для RabbitMQ
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(json.Msg),
		}

		// Публикуем сообщение в очередь
		if err := channelRabbitMQ.Publish(
			"",              // exchange
			"QueueService1", // queue name
			false,           // mandatory
			false,           // immediate
			message,
		); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Успешный ответ
		ctx.JSON(http.StatusOK, gin.H{"status": "message sent"})
	})

	r.Run(":8080")
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
	version  = "1.0.0"
)
