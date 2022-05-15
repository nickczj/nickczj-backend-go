package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickczj/web1/auth"
	"github.com/nickczj/web1/cache"
	"github.com/nickczj/web1/config"
	"github.com/nickczj/web1/database"
	log "github.com/sirupsen/logrus"
)

//go:generate go env -w GO111MODULE=on
//go:generate go mod tidy
//go:generate go mod download

var app *gin.Engine

func main() {
	config.SetEnvironment()

	if !gin.IsDebugging() {
		log.SetLevel(log.InfoLevel)
	}

	app = gin.Default()
	app.Use(config.Cors())
	app.Use(auth.Jwt())

	database.Init()
	cache.Init()
	initializeRoutes()

	app.Run(":8080")
}
