package configuration

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/inconshreveable/log15.v2"
)

func GlobalSecurity(context *gin.Context) {
	log15.Info("Global security is executed!")
	context.Next()
}
