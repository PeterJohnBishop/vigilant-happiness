package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func ServeGin() {
	r := gin.Default()

	AddPublicRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on :%s\n", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
