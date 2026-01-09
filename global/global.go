package global

import (
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	Client   *resty.Client
)
