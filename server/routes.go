package server

import "github.com/gin-gonic/gin"

func AddPublicRoutes(r *gin.Engine) {
	r.GET("/", Home())
	r.POST("/webhook/struct", HandleWebhookPayloadTypeMap())
	r.POST("/struct", HandleTypeMap())
	r.POST("/webhook/interface", HandleWebhookPayloadInterfaceMap())
	r.POST("/interface", HandleInterfaceMap())
}
