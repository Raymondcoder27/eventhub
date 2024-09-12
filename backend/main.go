package main

import "github.com/gin-gonic/gin"

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST("/createEvent", controllers.CreateEvent)
	r.GET("/events", controllers.EventsIndex)
	r.PUT("/events", controllers.UpdateEvent)
	r.DELETE("/event/:id", controllers.DeleteEvent)
	r.Run()
}
