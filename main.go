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
	restBoiler := controller.RestBoiler(database, cache, &appConfig) //dependency injection ke setiap implementor
	apiKeyInternal := appConfig.GetString("api.key.internal")

	auth := route.Group("/auth")
	auth.Use(configuration.ThirdPartySecurity(apiKeyInternal, cache))
	auth.POST("/v1/sign-up", restBoiler.CustomerSignup)
	auth.POST("/v1/sign-in", restBoiler.CustomerSignin)
	auth.POST("/v1/request-activation", restBoiler.CustomerResendActivation)

	data := route.Group("/data")
	data.Use(configuration.ThirdPartySecurity(apiKeyInternal, cache))
	data.GET("/v1/detail/:id", restBoiler.CustomerGetDetail)

	public := route.Group("/public")
	public.POST("/v1/activate", restBoiler.CustomerActivation)

	info := route.Group("/info")
	info.Use(configuration.GlobalSecurity(apiKeyInternal, cache, appConfig.GetString("jwt.secret")))
	info.GET("/v1/profile", restBoiler.CustomerGetProfile)

	route.Run(viper.GetString("application.port"))
}
