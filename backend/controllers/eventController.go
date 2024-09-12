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
	currentTime := time.Now()
	c.JSON(http.StatusOK, gin.H{"code": 200, "timestamp": currentTime})
}

// update event
func UpdateEvent(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")
	//get the data off the req body
	var body struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Location     string `json:"location"`
		Date         string `json:"date"`
		Category     string `json:"category"`
		MaxAttendees string `json:"maxAttendees"`
	}
	c.Bind(&body)

	//find the event you are updating
	var event models.Event
	initializers.DB.First(&event, id)

	//update it
	initializers.DB.Model(&event).Updates(models.Event{
		Name:         body.Name,
		Description:  body.Description,
		Location:     body.Location,
		Date:         body.Date,
		Category:     body.Category,
		MaxAttendees: body.MaxAttendees,
	})

	currentTime := time.Now()
	//respond with it
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": event, "timestamp": currentTime})
}

// delete event
func DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Event{}, id)
	// if err := initializers.DB.Delete(&event).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"messge": "Failed to delete event from database"})
	// 	return
	// }
	currentTime := time.Now()
	c.JSON(http.StatusOK, gin.H{"code": 200, "timestamp": currentTime})
}
