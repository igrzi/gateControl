package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/cadastro/initializers"
	"github.com/igrzi/gateControl/cadastro/models"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {
	userCPF, err := strconv.Atoi(c.Query("cpf"))
	userName := c.Query("name")
	userCategory := c.Query("category")

	if err != nil {
		c.JSON(400, "cpf parameter can't be empty and must be a number!")
		return
	}

	if checkIfCPFisNotRegistered(userCPF) {
		// Se o CPF não estiver registrado, cria o usuário
		user := models.User{Cpf: userCPF, Name: userName, Category: userCategory}
		initializers.DB.Create(&user)
		initializers.DB.Table("users").Where("cpf = ?", userCPF).Update("deleted_at", nil)

		c.JSON(200, "User created successfully!")
	} else {
		c.JSON(409, "User already on database")
	}
}

func UserShow(c *gin.Context) {
	var users []models.User

	// Query the database to get all users
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Create a slice to hold the user details
	var userDetails []gin.H
	for _, user := range users {
		userDetails = append(userDetails, gin.H{
			"name":     user.Name,
			"cpf":      user.Cpf,
			"category": user.Category,
		})
	}

	// Return the list of user details as JSON with specified field order
	c.JSON(200, gin.H{"users": userDetails})
}

// UserUpdate is a controller that updates a user CATEGORY in the database
func UserUpdate(c *gin.Context) {
	userCPF, err := strconv.Atoi(c.Query("cpf"))
	userCategory := c.Query("category")

	if err != nil {
		c.JSON(400, "cpf parameter can't be empty and must be a number!")
		return
	}

	if userCategory == "" {
		c.JSON(400, "category parameters can't be empty!")
	}

	if !checkIfCPFisNotRegistered(userCPF) {
		// Se o CPF estiver registrado, faz o update da categoria do usuário
		initializers.DB.Table("users").Where("cpf = ?", userCPF).Update("category", userCategory)

		c.JSON(200, "User updated successfully!")
	} else {
		c.JSON(409, "User not found on our database")
	}
}

func UserDelete(c *gin.Context) {
	userCPF, err := strconv.Atoi(c.Query("cpf"))

	if err != nil {
		c.JSON(400, "cpf parameter can't be empty and must be a number!")
		return
	}

	if !checkIfCPFisNotRegistered(userCPF) {
		// Se o CPF estiver registrado, faz o update da categoria do usuário
		initializers.DB.Table("users").Where("cpf = ?", userCPF).Delete(&models.User{})

		c.JSON(200, "User deleted successfully!")
	} else {
		c.JSON(409, "User not found on our database")
	}
}

func checkIfCPFisNotRegistered(cpf int) bool {
	var user models.User

	result := initializers.DB.Table("users").Where("cpf = ?", cpf).First(&user)

	// If the user is not found in the database, return true
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	// Log other possible errors and return false
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}

	return false
}
