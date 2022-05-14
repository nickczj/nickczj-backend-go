package global

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	//Firestore *firestore.Client
	Cache  *cache.Cache
	Client *resty.Client
)
