package main

import (
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"net/http"
)

func initializeRoutes() {
	// Metrics
	app.Use(stats.RequestStats())
	app.GET("/stats", getStats)

	// Finances
	privateRoutes := app.Group("/p/v1")
	{
		privateRoutes.GET("/nw/:id", GetNetWorth)
		privateRoutes.POST("/nw", SaveNetWorth)
	}

	app.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	})
}
