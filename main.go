package main

import (
	"morse/message"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/message", message.GetMessages)
	router.POST("/message", message.SendMessage)
	router.PUT("/message", message.EditMessage)

	router.Run("localhost:8080")
}
