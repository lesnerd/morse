package main

import (
	"morse/handlers/message"
	"morse/handlers/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/message", message.GetMessages)
	router.POST("/message", message.SendMessage)
	router.PUT("/message", message.EditMessage)

	router.GET("/user", user.GetUser)
	router.POST("/user", user.Register)
	router.PUT("/user", user.EditUser)
	router.PUT("/user/online", user.UserOnline)

	router.Run("localhost:8080")
}
