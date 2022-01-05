package configuration

import (
	"net/http"
	"omni-customer/utility"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/inconshreveable/log15.v2"
)

func GlobalSecurity(cache *redis.Client, tokenSecret string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		log15.Info("Global security is executed! ", tokenString)
		if result := utility.ValidateToken(cache, tokenString, tokenSecret); !result {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		context.Next()
	}

}
