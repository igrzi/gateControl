package main

import (
	"github.com/igrzi/gateControl/vagas/initializers"
	"github.com/igrzi/gateControl/vagas/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Spots{})
}
