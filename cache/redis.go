package cache

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/nickczj/web1/config"
	"github.com/nickczj/web1/global"
	"github.com/nickczj/web1/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func Init() {
	password, err := config.AccessSecretVersion("projects/171134391294/secrets/redis_password/versions/latest")
	if err != nil {
		log.Error("Error getting GCP secret ", err)
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: *password,
		DB:       0, // use default DB
	})

	global.Cache = cache.New(&cache.Options{
		Redis: rdb,
		//LocalCache: cache.NewTinyLFU(1000, time.Second*10),
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Error("Error initializing redis server ", err)
	} else {
		log.Info("Initialized redis server: ", pong)
	}
}

func GetElse(key string, f func() (model.Finances, error)) (model.Finances, error) {
	if global.Cache == nil {
		return f()
	}

	result := new(model.Finances)

	err := global.Cache.Once(&cache.Item{
		Key:   key,
		Value: result, // destination
		Do: func(*cache.Item) (interface{}, error) {
			log.Info("cache missed: ", key)
			return f()
		},
	})

	if err != nil {
		log.Error("Failed to load from cache: ", key, err)
		return f()
	}

	return *result, nil
}

func GenerateKey(args ...string) string {
	return strings.Join(args, "-")
}
