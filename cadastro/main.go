package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/cadastro/controllers"
	"github.com/igrzi/gateControl/cadastro/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// localhost:8020
	router.POST("/usuario/create", controllers.UserCreate)   // Cria um usuário
	router.GET("/usuario", controllers.UserShow)             // Mostra os usuários
	router.PUT("/usuario/update", controllers.UserUpdate)    // Atualiza um usuário
	router.DELETE("/usuario/delete", controllers.UserDelete) // Deleta um usuário

	// This will run the server on localhost:8020
	router.Run()
}
