package controller

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigurationHandler struct {
	DB        *gorm.DB
	Cache     *redis.Client
	AppConfig *viper.Viper
}

func RestBoiler(database *gorm.DB, cache *redis.Client, appConfig *viper.Viper) ConfigurationHandler {
	return ConfigurationHandler{database, cache, appConfig}
}
