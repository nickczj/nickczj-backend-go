package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickczj/web1/cache"
	"github.com/nickczj/web1/config"
	log "github.com/sirupsen/logrus"
)

var app *gin.Engine

func main() {
	config.SetEnvironment()

	if !gin.IsDebugging() {
		log.SetLevel(log.InfoLevel)
	}

	app = gin.Default()
	app.Use(config.Cors())

	//database.Init()
	cache.Init()
	initializeRoutes()

	app.Run(":8888")
}
