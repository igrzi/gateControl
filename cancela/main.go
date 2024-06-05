package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/cancela/controllers"
	"github.com/igrzi/gateControl/cancela/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	// localhost:8050
	router.POST("/cancela/open", controllers.BarrierOpen)   // Abre a cancela
	router.POST("/cancela/close", controllers.BarrierClose) // Fecha a cancela

	// This will run the server on localhost:8050
	router.Run()
}
