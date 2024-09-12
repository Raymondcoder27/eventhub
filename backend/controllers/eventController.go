package controllers

import (
	"net/http"
	"time"

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
		Name string
		Description string
		Location string
		Date   time.Time
		Category string
		MaxAttendees string
	}
	c.Bind(&eventRequest)

	err := initializers.DB.Create(&eventRequest{
		Name:         eventRequest.Name,
		Description:  eventRequest.Description,
		Location:     eventRequest.Description,
		Date:         eventRequest.Date,
		Category:     eventRequest.Category,
		MaxAttendees: eventRequest.MaxAttendees,
		CreatedAt:    time.Now(),
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving document metadata in database: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200})
}
