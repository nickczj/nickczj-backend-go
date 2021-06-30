package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getNetWorth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"netWorth": calculateNetWorth()})
}

func calculateNetWorth() float64 {
	return 170000
}
