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
		// @Summary get net worth
		finances.GET("/nw/:id", GetNetWorth)
		// @Summary save net worth
		finances.POST("/nw", SaveNetWorth)
	}

	// @Summary get current PSI & UV Index from data.gov.sg
	app.GET("/weather/now", GetWeatherNow)

	// @Summary get current invalign tray I'm on
	app.GET("/invisalign/current-tray", GetCurrentTray)

	app.GET("/flights", GetFlights)
	app.GET("/flights/multi", GetFlightsMulti)

	// Metrics
	app.Use(stats.RequestStats())
	app.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}
