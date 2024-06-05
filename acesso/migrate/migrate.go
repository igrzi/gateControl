package main

import (
	"github.com/igrzi/gateControl/acesso/initializers"
	"github.com/igrzi/gateControl/acesso/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Access{})
}
