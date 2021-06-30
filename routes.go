package main

import (
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"net/http"
)

func initializeRoutes() {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	})

	router.Use(stats.RequestStats())
	router.GET("/stats", getStats)

	privateRoutes := router.Group("/p/v1")
	{
		privateRoutes.GET("/nw", getNetWorth)
	}
}
