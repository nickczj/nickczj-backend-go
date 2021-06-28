package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/personal/netWorth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"netWorth": 170000})
	})

	router.Run(":8888")
}
