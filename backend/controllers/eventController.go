package controllers

import (
	"net/http"

	"github.com/eventhub/models"
	"github.com/gin-gonic/gin"
)

// getting all events
func EventsIndex(c *gin.Context) {
	var events []models.Event
	c.JSON(http.StatusOK, gin.H{"message": events})
}
