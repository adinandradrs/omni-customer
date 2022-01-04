package main

import (
	"omni-customer/configuration"
	"omni-customer/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	database := configuration.ConfigDatabase()
	restRepository := controller.RestRepository(database)
	route.POST("/v1/sign-up", restRepository.CustomerSignup)
	route.POST("/v1/sign-in", restRepository.CustomerSignin)
	route.POST("/v1/activate", restRepository.CustomerActivation)
	route.Run()
}
