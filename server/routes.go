package server

import "github.com/gin-gonic/gin"

func AddPublicRoutes(r *gin.Engine) {
	r.GET("/", Home())
	r.POST("/webhook", HandleWebhookPayload())
	r.POST("/webhook/struct", HandleWebhookPayloadStruct())
}
