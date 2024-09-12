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
	var eventRequest models.Event
	c.Bind(&eventRequest)

	initializers.DB.Create(models.Event{
		Name:         eventRequest.Name,
		Description:  eventRequest.Description,
		Location:     eventRequest.Description,
		Date:         eventRequest.Date,
		Category:     eventRequest.Category,
		MaxAttendees: eventRequest.MaxAttendees,
		CreatedAt:    time.Now(),
	})
}
