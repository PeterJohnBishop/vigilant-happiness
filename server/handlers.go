package server

import (
	"net/http"
	"strings"
	"vigilant-happiness/server/services"

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
		name := payload["event"].(string)
		if name == "" {
			name = "WebhookEvent"
		}
		output, err := services.GenerateTypeMap(name, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate type map",
				"data":  err.Error(),
			})
			return
		}

		clean := strings.ReplaceAll(output, `\n`, "\n")
		clean = strings.ReplaceAll(clean, `\"`, `"`)
		clean = strings.TrimPrefix(clean, "```go")
		clean = strings.TrimPrefix(clean, "```")
		clean = strings.TrimSuffix(clean, "```")
		clean = strings.TrimSpace(clean)

		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(clean))
	}
}

func HandleTypeMap() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload map[string]interface{}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON payload",
				"data":  err.Error(),
			})
			return
		}
		name := c.Query("name")
		if name == "" {
			name = "GeneratedType"
		}
		output, err := services.GenerateTypeMap(name, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate type map",
				"data":  err.Error(),
			})
			return
		}

		clean := strings.ReplaceAll(output, `\n`, "\n")
		clean = strings.ReplaceAll(clean, `\"`, `"`)
		clean = strings.TrimPrefix(clean, "```go")
		clean = strings.TrimPrefix(clean, "```")
		clean = strings.TrimSuffix(clean, "```")
		clean = strings.TrimSpace(clean)

		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(clean))
	}
}
