package controller

import (
	"net/http"
	"omni-customer/model/request"
	"omni-customer/model/response"

	"github.com/gin-gonic/gin"
)

func (repository DatabaseHandler) CustomerSignin(context *gin.Context) {
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
