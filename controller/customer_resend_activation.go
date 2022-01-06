package controller

import (
	"net/http"
	"omni-customer/model/entity"
	"omni-customer/model/request"
	"omni-customer/model/response"
	constants "omni-customer/utility"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/inconshreveable/log15.v2"
)

func (boiler ConfigurationHandler) CustomerResendActivation(context *gin.Context) {
	var input request.CustomerEmailRequest
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
	var existingCustomer entity.Customer
	if result := boiler.DB.
		Where("is_deleted = false").
		Where("status = ?", constants.CUSTOMER_STATUS_REGISTERED).
		Where("activation_id is null").
		Where("email = ?", input.Email).First(&existingCustomer); result.Error != nil {
		log15.Error("Customer activation failed to find with error = ", result.Error)
		context.JSON(http.StatusNotFound, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_DATA_NOT_FOUND,
			Result:  false,
		})
		return
	}
	dataCacheActivationId, errorRedis1 := boiler.Cache.Get(constants.CACHE_CUSTOMER_ACTIVATION_EMAIL + existingCustomer.Email.String).Result()
	boiler.Cache.Del(constants.CACHE_CUSTOMER_ACTIVATION + dataCacheActivationId)
	boiler.Cache.Del(constants.CACHE_CUSTOMER_ACTIVATION_EMAIL + existingCustomer.Email.String)
	uid := uuid.New()
	errorRedis2 := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_ACTIVATION+uid.String(), existingCustomer.Email.String, boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
	errorRedis3 := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_ACTIVATION_EMAIL+existingCustomer.Email.String, uid.String(), boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
	if errorRedis1 == nil || errorRedis2 == nil || errorRedis3 == nil {
		context.JSON(http.StatusOK, response.BaseResponse{
			Data:    nil,
			Message: constants.SUCCESS_MSG_DATA_SUBMIT,
			Result:  true,
		})
		return
	}
	context.JSON(http.StatusForbidden, response.BaseResponse{
		Data:    nil,
		Message: constants.ERR_MSG_UNAUTHORIZED,
		Result:  false,
	})
}
