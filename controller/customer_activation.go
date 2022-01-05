package controller

import (
	"database/sql"
	"net/http"
	"omni-customer/model/entity"
	"omni-customer/model/request"
	"omni-customer/model/response"
	constants "omni-customer/utility"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/inconshreveable/log15.v2"
)

func (boiler ConfigurationHandler) CustomerActivation(context *gin.Context) { //dependency injection boiler
	var input request.CustomerActivationRequest
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
	dataCache, error := boiler.Cache.Get(constants.CACHE_CUSTOMER_ACTIVATION + input.ActivationId).Result()
	if error == redis.Nil {
		log15.Error("Redis is empty for activation with activationId = ", input.ActivationId)
		context.JSON(http.StatusUnauthorized, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_UNAUTHORIZED,
			Result:  false,
		})
		return
	} else if error != nil {
		log15.Error("Redis is empty for activation with error = ", error.Error())
		context.JSON(http.StatusInternalServerError, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_SOMETHING_WENT_WRONG,
			Result:  false,
		})
		return
	} else {
		var existingCustomer entity.Customer
		if result := boiler.DB.Where("activation_id = ?", input.ActivationId).
			Where("is_deleted = false").
			Where("status = ?", constants.CUSTOMER_STATUS_REGISTERED).Where("email = ?", dataCache).First(&existingCustomer); result.Error != nil {
			log15.Error("Customer activation failed to find with error = ", result.Error)
			context.JSON(http.StatusNotFound, response.BaseResponse{
				Data:    nil,
				Message: constants.ERR_MSG_DATA_NOT_FOUND,
				Result:  false,
			})
			return
		}
		existingCustomer.Status = constants.CUSTOMER_STATUS_ACTIVATED
		existingCustomer.ActivationDate = sql.NullTime{Time: time.Now(), Valid: true}
		boiler.DB.Save(existingCustomer)
		boiler.Cache.Del(constants.CACHE_CUSTOMER_ACTIVATION + input.ActivationId)
		context.JSON(http.StatusOK, response.BaseResponse{
			Data:    nil,
			Message: constants.SUCCESS_MSG_DATA_SUBMIT,
			Result:  true,
		})
	}

}
