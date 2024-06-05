package main

import (
	"github.com/igrzi/gateControl/creditos/initializers"
	"github.com/igrzi/gateControl/creditos/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Credits{})
}
