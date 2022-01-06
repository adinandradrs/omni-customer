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
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/inconshreveable/log15.v2"
)

func (boiler ConfigurationHandler) CustomerSignup(context *gin.Context) { //dependency injection boiler
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
	encryptedPassword, error := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if error != nil {
		context.JSON(http.StatusInternalServerError, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
	customer := entity.Customer{
		Fullname:    sql.NullString{String: input.Fullname, Valid: true},
		Email:       sql.NullString{String: input.Email, Valid: true},
		Password:    sql.NullString{String: string(encryptedPassword), Valid: true},
		Status:      constants.CUSTOMER_STATUS_REGISTERED,
		Address:     sql.NullString{String: input.Address, Valid: true},
		PhoneNo:     sql.NullString{String: input.PhoneNo, Valid: true},
		IsDeleted:   false,
		CreatedDate: time.Now(),
	}
	if error := boiler.DB.Create(&customer).Error; error != nil {
		log15.Error("Failed to create an user with email = ", input.Email)
		context.JSON(http.StatusInternalServerError, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_SOMETHING_WENT_WRONG,
			Result:  false,
		})
	} else {
		errorRedis1 := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_ACTIVATION+uid.String(), customer.Email.String, boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
		errorRedis2 := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_ACTIVATION_EMAIL+customer.Email.String, uid.String(), boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
		if errorRedis1 != nil || errorRedis2 != nil {
			panic("Failed to set activation cache on signup")
		}

		log15.Info("Customer successfuly registered with email = ", input.Email)
		context.JSON(http.StatusOK, response.BaseResponse{
			Data:    nil,
			Message: constants.SUCCESS_MSG_DATA_SUBMIT,
			Result:  true,
		})
	}
}
