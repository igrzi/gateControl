package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/vagas/controllers"
	"github.com/igrzi/gateControl/vagas/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// localhost:8040
	router.GET("/vagas", controllers.ShowSpots)                // Mostra as vagas
	router.POST("/vagas/adjust", controllers.AdjustAmountSpot) // Cria uma vaga
	router.POST("/vagas/occuppy", controllers.OccuppySpot)     // Ocupa uma vaga
	router.POST("/vagas/vacate", controllers.VacateSpot)       // Libera uma vaga

	// This will run the server on localhost:8040
	router.Run()
}
