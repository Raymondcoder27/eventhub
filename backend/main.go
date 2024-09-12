package main

import (
	"fmt"

	"github.com/eventhub/controllers"
	"github.com/eventhub/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.POST("/createEvent", controllers.CreateEvent)
	r.GET("/events", controllers.EventsIndex)
	// r.PUT("/events", controllers.UpdateEvent)
	// r.DELETE("/event/:id", controllers.DeleteEvent)
	fmt.Println("Hello World")
	r.Run()
}
