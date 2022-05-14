package database

import (
	"fmt"
	"github.com/nickczj/web1/config"
	"github.com/nickczj/web1/global"
	"github.com/nickczj/web1/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewMariaDB() {
	handler := dbHandler(start)
	dsn := "nczj:nich0laS@tcp(nczj-web1-mariadb:3306)/nick1?charset=utf8mb4&parseTime=True&loc=Local"
	handler.handleDb(mysql.Open(dsn))
}

func NewPostgres() {
	handler := dbHandler(start)
	password, err := config.AccessSecretVersion("projects/171134391294/secrets/db_password/versions/latest")
	if err != nil {
		log.Error("Error getting GCP secret ", err)
		return
	}
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=verify-full",
		viper.GetString("db.username"),
		*password,
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	handler.handleDb(postgres.Open(dsn))
}

func start(dialect gorm.Dialector) error {
	// initialize connection
	conf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	database, err := gorm.Open(dialect, conf)
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

type dbHandler func(dialect gorm.Dialector) error

func (fn dbHandler) handleDb(dialect gorm.Dialector) {
	if err := fn(dialect); err != nil {
		log.Error(err)
	}
}
