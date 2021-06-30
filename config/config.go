package config

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func SetEnvironment() {
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	var env string
	if os.Getenv("APP_ENV") != "" {
		env = strings.ToLower(os.Getenv("APP_ENV"))
	} else {
		env = "local"
	}
	log.Info("Application environment: ", env)

	viper.SetConfigName(env)
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to load config", err)
	}

	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}
