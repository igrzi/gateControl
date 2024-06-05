package main

import (
	"github.com/igrzi/gateControl/cadastro/initializers"
	"github.com/igrzi/gateControl/cadastro/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
