package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", func(c *gin.Context) {
		events := models.GetAllEvents()
		c.JSON(http.StatusOK, events)
	})

	server.POST("/events", func(c *gin.Context) {
		var event models.Event
		err := c.ShouldBindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "fill required fields",
			})
			return
		}
		event.ID = 1
		event.UserID = 1
		c.JSON(http.StatusOK, gin.H{
			"message": "event created",
			"event":   event,
		})
		event.Save()
	})

	server.Run(":8080")
}
