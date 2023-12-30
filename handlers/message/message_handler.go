package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Message represents a message sent by the user
type Message struct {
	FromId string `json:"from"`
	ToId   string `json:"to"`
	Text   string `json:"text"`
}

func GetMessages(c *gin.Context) {
	// Return the messages slice that we declared above
	c.JSON(http.StatusOK, gin.H{"messages": "messages"})
}

func SendMessage(c *gin.Context) {
	// Get the message from the request body
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the message to the in-memory slice
	//messages = append(messages, message)
	c.JSON(http.StatusCreated, gin.H{"status": msg})
}

func EditMessage(c *gin.Context) {
	// Get the message from the request body
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
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
