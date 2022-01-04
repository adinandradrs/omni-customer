package controller

import (
	"net/http"
	"omni-customer/model/request"
	"omni-customer/model/response"

	"github.com/gin-gonic/gin"
)

func (boiler ConfigurationHandler) CustomerSignin(context *gin.Context) { //dependency injection boiler
	var input request.CustomerSigninRequest
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
}
