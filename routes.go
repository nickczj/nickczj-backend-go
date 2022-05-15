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

	finances := app.Group("/p/v1")
	{
		finances.GET("/nw/:id", GetNetWorth)
		finances.POST("/nw", SaveNetWorth)
	}

	app.GET("/weather/now", GetWeatherNow)
	app.GET("/invisalign/current-tray", GetCurrentTray)

	// Metrics
	app.Use(stats.RequestStats())
	app.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}
