package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/acesso/controllers"
	"github.com/igrzi/gateControl/acesso/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// localhost:8010
	router.POST("/acesso/entry", controllers.AccessEntry)
	router.POST("/acesso/leave", controllers.AccessLeave)

	// This will run the server on localhost:8010
	router.Run()
}
