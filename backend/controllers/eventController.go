package controllers

import (
	"net/http"

	"github.com/eventhub/initializers"
	"github.com/eventhub/models"
	"github.com/gin-gonic/gin"
)

// getting all events
func EventsIndex(c *gin.Context) {
	var events []models.Event
	c.JSON(http.StatusOK, gin.H{"message": events})
}

// creating an event
func CreateEvent(c *gin.Context) {
	var eventRequest struct {
		Name        string
		Description string
		Location    string
		// Date         time.
		Category     string
		MaxAttendees string
	}
	c.Bind(&eventRequest)

	if err := initializers.DB.Create(&eventRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving event into database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200})
}
