package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/acesso/initializers"
	"github.com/igrzi/gateControl/acesso/models"
)

func AccessEntry(c *gin.Context) {
	cpf := c.Query("cpf")

	// Check spot availability
	err := checkSpotAvailability()
	if err != nil {
		c.JSON(500, gin.H{"error": "There isn't any spot available"})
		return
	}

	// Open gate
	openGate()
	occupySpot()

	// Log access
	err = registerAccess(cpf, "entrada")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to log access"})
		return
	}

	c.JSON(200, gin.H{"message": "Entry allowed"})
}

func AccessLeave(c *gin.Context) {
	cpf := c.Query("cpf")

	err := subtractCreditFromUser(cpf)
	if err != nil {
		c.JSON(500, gin.H{"error": "Not enough credits available"})
		return
	}

	// Open gate
	openGate()
	releaseSpot()

	// Log access
	err = registerAccess(cpf, "saida")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to log access"})
		return
	}

	c.JSON(200, gin.H{"message": "Leave allowed"})
}

func subtractCreditFromUser(cpf string) error {
	// Prepare the URL with query parameters
	apiUrl := "http://localhost:8030/creditos/use"
	params := url.Values{}
	params.Set("cpf", cpf)
	urlWithParams := apiUrl + "?" + params.Encode()

	// Send the HTTP POST request
	_, err := http.Post(urlWithParams, "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

func occupySpot() error {
	resp, err := http.Post("http://localhost:8040/vagas/occuppy", "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("failed to occupy spot: bad request")
	}

	return nil
}

func releaseSpot() error {
	resp, err := http.Post("http://localhost:8040/vagas/vacate", "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("failed to release spot: bad request")
	}

	return nil
}

func checkSpotAvailability() error {
	resp, err := http.Get("http://localhost:8040/vagas")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var availability map[string]int
	err = json.NewDecoder(resp.Body).Decode(&availability)
	if err != nil {
		return err
	}

	quantityAvailable, ok := availability["quantity_available"]
	if !ok {
		return errors.New("quantity_available not found in response")
	}

	if quantityAvailable <= 0 {
		return errors.New("no spots available")
	}

	return nil
}

func openGate() {
	// Call cancela microservice to open the gate
	_, _ = http.Post("http://localhost:8050/cancela/open", "application/json", nil)
}

func registerAccess(cpf, accessType string) error {
	convertedCpf, _ := strconv.Atoi(cpf)

	access := models.Access{Cpf: convertedCpf, Type: accessType}
	err := initializers.DB.Create(&access).Error
	if err != nil {
		return err
	}
	return nil
}
