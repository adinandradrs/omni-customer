package main

import (
	"omni-customer/configuration"
	"omni-customer/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	route := gin.Default()
	appConfig := *configuration.GetConfiguration()
	database := configuration.ConfigDatabase(appConfig.GetString("datasource.url"))
	cache := configuration.ConfigCache(appConfig.GetString("redis.host")+":"+appConfig.GetString("redis.port"), appConfig.GetString("redis.password"))
	restRepository := controller.RestBoiler(database, cache, &appConfig)
	route.POST("/v1/sign-up", restRepository.CustomerSignup)
	route.POST("/v1/sign-in", restRepository.CustomerSignin)
	route.POST("/v1/activate", restRepository.CustomerActivation)
	route.Run(viper.GetString("application.port"))
}
