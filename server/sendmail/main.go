package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	sm := router.Group("/api/v1")

	sm.GET("/send-mail/ping", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := router.Run(httpPort()); err != nil {
		log.Fatal(err)
	}
}

func httpPort() string {
	port := "6001"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	return fmt.Sprintf(":%s", port)
}
