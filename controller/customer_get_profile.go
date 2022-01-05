package controller

import (
	"net/http"
	"omni-customer/model/entity"
	"omni-customer/model/response"
	"omni-customer/utility"
	constants "omni-customer/utility"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gopkg.in/inconshreveable/log15.v2"
)

func (boiler ConfigurationHandler) CustomerGetProfile(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	customerLoginResponse, errorCustomerInfo := utility.GetCustomerInfo(boiler.Cache, tokenString, boiler.AppConfig.GetString("jwt.secret"))
	if errorCustomerInfo != nil {
		context.JSON(http.StatusUnauthorized, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_UNAUTHORIZED,
			Result:  false,
		})
	}
	var existingCustomer entity.Customer
	if result := boiler.DB.
		Where("email = ?", customerLoginResponse.Email).
		Where("is_deleted = false").
		First(&existingCustomer); result.Error != nil {
		log15.Error("Customer profile failed to find with error = ", result.Error)
		context.JSON(http.StatusNotFound, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_DATA_NOT_FOUND,
			Result:  false,
		})
		return
	}
	context.JSON(http.StatusOK, response.BaseResponse{
		Data:    structs.Map(response.CustomerDetail{Email: existingCustomer.Email, Fullname: existingCustomer.Fullname}),
		Message: constants.SUCCESS_MSG_DATA_FOUND,
		Result:  true,
	})
}
