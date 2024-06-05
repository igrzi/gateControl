package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/gateControl/vagas/initializers"
	"github.com/igrzi/gateControl/vagas/models"
)

func ShowSpots(c *gin.Context) {
	// Find the spots record
	var spots models.Spots
	result := initializers.DB.First(&spots)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Spots record not found"})
		return
	}

	c.JSON(200, gin.H{"quantity_available": spots.QuantityAvailable, "max_quantity": spots.MaxQuantity})
}

// AdjustAmountSpot adjusts the amount of spots available or creates a new spot record if none exists
func AdjustAmountSpot(c *gin.Context) {
	quantitySpots, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.JSON(400, gin.H{"error": "quantity parameter can't be empty and must be a number!"})
		return
	}

	// Check if there's a record in the database
	var spot models.Spots
	result := initializers.DB.First(&spot)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			// If no record is found, create a new record
			newSpot := models.Spots{QuantityAvailable: quantitySpots, MaxQuantity: quantitySpots}
			initializers.DB.Create(&newSpot)
			c.JSON(200, gin.H{"message": "New spot record created successfully!", "quantity_available": newSpot.QuantityAvailable, "max_quantity": newSpot.MaxQuantity})
		} else {
			// Handle other potential errors
			c.JSON(500, gin.H{"Internal error": result.Error.Error()})
		}
		return
	}

	// If a record is found, update the max quantity and adjust the available spots accordingly
	spot.MaxQuantity = quantitySpots
	if spot.QuantityAvailable > spot.MaxQuantity {
		spot.QuantityAvailable = spot.MaxQuantity
	}
	initializers.DB.Save(&spot)
	c.JSON(200, gin.H{"message": "Spot record updated successfully!", "quantity_available": spot.QuantityAvailable, "max_quantity": spot.MaxQuantity})
}

func OccuppySpot(c *gin.Context) {
	// Find the spots record
	var spots models.Spots
	result := initializers.DB.First(&spots)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Spots record not found"})
		return
	}

	// Decrease the QuantityAvailable by 1 if available
	if spots.QuantityAvailable > 0 {
		spots.QuantityAvailable -= 1
	} else {
		c.JSON(400, gin.H{"error": "All spots are currently occupied"})
		return
	}

	// Save the updated spots back to the database
	initializers.DB.Save(&spots)

	c.JSON(200, gin.H{"message": "One spot used!", "available_spots": spots.QuantityAvailable})
}

func VacateSpot(c *gin.Context) {
	// Find the spots record
	var spots models.Spots
	result := initializers.DB.First(&spots)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Spots record not found"})
		return
	}

	// Increase the QuantityAvailable by 1 if it's less than MaxQuantity
	if spots.QuantityAvailable < spots.MaxQuantity {
		spots.QuantityAvailable += 1
	} else {
		c.JSON(400, gin.H{"error": "All spots are currently available"})
		return
	}

	// Save the updated spots back to the database
	initializers.DB.Save(&spots)

	c.JSON(200, gin.H{"message": "One spot vacated!", "available_spots": spots.QuantityAvailable})
}
