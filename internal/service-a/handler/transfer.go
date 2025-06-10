package handler

import (
	"encoding/json"
	"net/http"
	"transaction-worker/internal/common/rabbitmq"
	"transaction-worker/internal/service-b/model"

	"github.com/gin-gonic/gin"
)

func TransferHandler(rabbit *rabbitmq.RabbitMQ) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.TransferRequest

		if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := user.ValidateInput(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Вызов сервиса
		// err :=

		data, err := json.Marshal(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal user"})
			return
		}

		if err := rabbit.Publish("QueueService1", []byte(data)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "message sent"})
	}
}
