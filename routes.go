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
	finances := app.Group("/p/v1")
	{
		finances.GET("/nw/:id", GetNetWorth)
		finances.POST("/nw", SaveNetWorth)
	}

	// Weather
	weather := app.Group("/weather")
	{
		weather.GET("/now", GetWeatherNow)
	}

	// Invisalign
	invisalign := app.Group("/invisalign")
	{
		invisalign.GET("/current-tray", GetCurrentTray)
	}

	app.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hi": "there"})
	})
}
