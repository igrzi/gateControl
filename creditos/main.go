package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/creditos/controllers"
	"github.com/igrzi/gateControl/creditos/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// localhost:8030
	router.POST("/creditos/add", controllers.AddCredits) // Adiciona créditos
	router.POST("/creditos/use", controllers.UseCredits) // Usa créditos

	// This will run the server on localhost:8030
	router.Run()
}
