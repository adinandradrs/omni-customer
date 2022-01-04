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

func (boiler ConfigurationHandler) CustomerSignup(context *gin.Context) {
	var input request.CustomerSignupRequest
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
	uid := uuid.New()
	customer := entity.Customer{
		Fullname:     input.Fullname,
		Email:        input.Email,
		Password:     input.Password,
		Status:       constants.CUSTOMER_STATUS_REGISTERED,
		IsDeleted:    false,
		CreatedDate:  time.Now(),
		ActivationId: uid.String(),
	}
	if error := boiler.DB.Create(&customer).Error; error != nil {
		log15.Error("Failed to create an user with email = ", input.Email)
		context.JSON(http.StatusInternalServerError, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_SOMETHING_WENT_WRONG,
			Result:  false,
		})
	} else {
		err := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_ACTIVATION+uid.String(), customer.Email, boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
		if err != nil {
			panic(err)
		}

		log15.Info("Customer successfuly registered with email = ", input.Email)
		context.JSON(http.StatusOK, response.BaseResponse{
			Data:    nil,
			Message: constants.SUCCESS_MSG_DATA_SUBMIT,
			Result:  true,
		})
	}
}
