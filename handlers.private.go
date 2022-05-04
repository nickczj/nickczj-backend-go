package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickczj/web1/model"
	"github.com/nickczj/web1/service"
	"net/http"
	"strconv"
)

func GetNetWorth(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"bad request": "invalid id provided"})
	} else if finances, err := service.GetNetWorth(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve net worth"})
	} else {
		c.JSON(http.StatusOK, gin.H{"netWorth": finances.NetWorth})
	}
}

func SaveNetWorth(c *gin.Context) {
	var finances model.Finances
	_ = c.ShouldBindJSON(&finances)
	if f, err := service.SaveNetWorth(finances); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unnable to save net worth"})
	} else {
		c.JSON(http.StatusOK, gin.H{"netWorth": f.NetWorth})
	}
}

func GetWeatherNow(c *gin.Context) {
	if weather, err := service.Now(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to retrieve weather"})
	} else {
		c.JSON(http.StatusOK, gin.H{"weather": weather})
	}
}

func GetCurrentTray(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tray_number": service.CurrentTray()})
}
