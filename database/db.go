package database

import (
	"github.com/nickczj/web1/global"
	"github.com/nickczj/web1/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Init() {
	handler := dbHandler(start)
	handler.handleDb()
}

func start() error {
	// initialize connection
	dsn := "nczj:nich0laS@tcp(nczj-web1-mariadb:3306)/nick1?charset=utf8mb4&parseTime=True&loc=Local"
	conf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	database, err := gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		return err
	}

	// settings
	sqlDB, err := database.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)

	// test connection
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	// synchronize DB schemas
	err = database.AutoMigrate(&model.Finances{})
	if err != nil {
		return err
	}

	global.Database = database
	return nil
}

type dbHandler func() error

func (fn dbHandler) handleDb() {
	if err := fn(); err != nil {
		log.Error(err)
	}
}
