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
	if err := initializers.DB.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching events."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": events})
}

// creating an event
func CreateEvent(c *gin.Context) {
	var eventRequest models.Event
	c.Bind(&eventRequest)

	if err := initializers.DB.Create(&eventRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error saving event into database."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200})
}

// update event
func UpdateEvent(c *gin.Context) {

}

// delete event
func DeleteEvent(c *gin.Context) {
	var event models.Event
	if err := initializers.DB.Delete(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messge": "Failed to delete event from database"})
	}
}
