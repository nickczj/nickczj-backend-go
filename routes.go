package main

import (
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"net/http"
)

func initializeRoutes() {
	app.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	})

	// Finances
	privateRoutes := app.Group("/p/v1")
	{
		privateRoutes.GET("/nw", getNetWorth)
	}

	// Metrics
	app.Use(stats.RequestStats())
	app.GET("/stats", getStats)
}
