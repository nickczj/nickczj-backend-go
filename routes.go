package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes() {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	})

	privateRoutes := router.Group("/p/v1")
	{
		privateRoutes.GET("/nw", getNetWorth)
	}
}
