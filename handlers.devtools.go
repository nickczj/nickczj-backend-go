package main

import (
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"net/http"
)

func getStats(c * gin.Context) {
	c.JSON(http.StatusOK, stats.Report())
}
