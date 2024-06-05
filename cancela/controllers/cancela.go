package controllers

import (
	"github.com/gin-gonic/gin"
)

func BarrierOpen(c *gin.Context) {

	c.JSON(200, "Barrier opened successfully!")
}

func BarrierClose(c *gin.Context) {

	c.JSON(200, "Barrier closed successfully!")
}
