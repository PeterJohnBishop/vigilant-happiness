package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Hello from Vigilnt-Happiness!",
		})
	}
}

func HandleWebhookPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload map[string]interface{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON payload",
				"data":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Payload received successfully",
			"data":    payload,
		})
	}
}
