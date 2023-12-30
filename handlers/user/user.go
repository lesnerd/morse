package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	UserOnline bool   `json:"user_online"`
	Connection string `json:"connection"` // will be used as a direct connection to user
}

func Register(c *gin.Context) {
	// Get the message from the request body
	var usr User
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the message to the in-memory slice
	//messages = append(messages, message)
	c.JSON(http.StatusCreated, gin.H{"status": usr})
}

func EditUser(c *gin.Context) {
	// Get the message from the request body
	var usr User
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the message ID from the URL
	//messageID := c.Param("id")

	// Find the message and update it
	// for i := 0; i < len(messages); i++ {
	// 	if messages[i].ID == messageID {
	// 		messages[i].Text = message.Text
	// 		break
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetUser(c *gin.Context) {
	// Return the messages slice that we declared above
	c.JSON(http.StatusOK, gin.H{"users": "users"})
}

func UserOnline(c *gin.Context) {
	// Get the message from the request body
	var usr User
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the message ID from the URL
	//messageID := c.Param("id")

	// Find the message and update it
	// for i := 0; i < len(messages); i++ {
	// 	if messages[i].ID == messageID {
	// 		messages[i].Text = message.Text
	// 		break
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
