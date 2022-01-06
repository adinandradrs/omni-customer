package configuration

import (
	"net/http"
	"omni-customer/utility"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/inconshreveable/log15.v2"
)

func GlobalSecurity(internalApiKey string, cache *redis.Client, tokenSecret string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := utility.GetBearerToken(context)
		apiKey := utility.GetApiKey(context)
		log15.Info("Global security is executed! ", tokenString)
		if result := utility.ValidateToken(cache, tokenString, tokenSecret); !result {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		doRunValidate(context, apiKey, internalApiKey, cache)
	}

}

func validateIsUsingInternalApiKey(apiKey string, internalApiKey string) bool {
	return apiKey == internalApiKey
}

func doRunValidate(context *gin.Context, apiKey string, internalApiKey string, cache *redis.Client) {
	if len(apiKey) < 1 {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
	if validateIsUsingInternalApiKey(apiKey, internalApiKey) {
		context.Next()
	} else {
		log15.Error("API Key is not defined yet for security passport")
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
