package main

import "github.com/gin-gonic/gin"

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	r.POST()
	r.GET()
	r.PUT()
	r.DELETE()
	r.Run()
}
