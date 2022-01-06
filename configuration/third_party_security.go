package configuration

import (
	"omni-customer/utility"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/inconshreveable/log15.v2"
)

func ThirdPartySecurity(internalApiKey string, cache *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKey := utility.GetApiKey(context)
		log15.Info("Third party security is executed! ", apiKey)
		doRunValidate(context, apiKey, internalApiKey, cache)
	}

}
