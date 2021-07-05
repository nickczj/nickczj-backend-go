package database

import (
	"github.com/nickczj/web1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func Init() {
	var err error
	// initialize connection
	dsn := "nczj:nich0laS@tcp(nczj-web1-mariadb:3306)/nick1?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	handleError(err)

	// settings
	sqlDB, err := database.DB()
	handleError(err)
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)

	// test connection
	err = sqlDB.Ping()
	handleError(err)

	// synchronize DB schemas
	err = database.AutoMigrate(&model.Finances{})
	if err != nil {
		log.Error("Failed to synchronize schemas. ", err)
	}

	DB = database
}

func handleError(err error) {
	if err != nil {
		log.Error("Failed to open database connection. ", err)
	}
}
