package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/inconshreveable/log15.v2"
)

func ThirdPartySecurity(context *gin.Context) {
	apiKey := context.GetHeader("x-api-key")
	log15.Info("Third party security is executed! ", apiKey)
	if len(apiKey) < 1 {
		context.AbortWithStatus(http.StatusUnauthorized)
	}
	context.Next()
}
