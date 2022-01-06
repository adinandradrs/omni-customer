package controller

import (
	"net/http"
	"omni-customer/model/response"
	"omni-customer/utility"
	constants "omni-customer/utility"

	"github.com/gin-gonic/gin"
)

func (boiler ConfigurationHandler) CustomerGetProfile(context *gin.Context) {
	tokenString := utility.GetBearerToken(context)
	customerLoginResponse, errorCustomerInfo := utility.GetCustomerInfo(boiler.Cache, tokenString, boiler.AppConfig.GetString("jwt.secret"))
	if errorCustomerInfo != nil {
		context.JSON(http.StatusUnauthorized, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_UNAUTHORIZED,
			Result:  false,
		})
	}
	getCustomerDetail(0, customerLoginResponse.Email, false, context, boiler)
}
