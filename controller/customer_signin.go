package controller

import (
	"encoding/json"
	"net/http"
	"omni-customer/model/entity"
	"omni-customer/model/request"
	"omni-customer/model/response"
	"omni-customer/utility"
	constants "omni-customer/utility"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/inconshreveable/log15.v2"
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
	var existingCustomer entity.Customer
	if result := boiler.DB.Select("id, password, email, fullname").Where("email = ?", input.Email).Where("status = ?", constants.CUSTOMER_STATUS_ACTIVATED).First(&existingCustomer); result.Error != nil {
		log15.Error("Customer sign in failed to find with error = ", result.Error)
		context.JSON(http.StatusUnauthorized, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_UNAUTHORIZED,
			Result:  false,
		})
		return
	}
	invalid := bcrypt.CompareHashAndPassword([]byte(existingCustomer.Password.String), []byte(input.Password))
	if invalid != nil {
		log15.Error("Customer status is forbidden with error, password mismatch")
		context.JSON(http.StatusForbidden, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_UNAUTHORIZED,
			Result:  false,
		})
		return
	} else {
		customerLoginResponse := response.CustomerLoginResponse{Email: existingCustomer.Email.String, UserId: int(existingCustomer.Id), Fullname: existingCustomer.Fullname.String}
		customerLoginResponse.Token, _ = utility.GenerateToken(existingCustomer.Id, &customerLoginResponse, boiler.AppConfig.GetUint("jwt.expiration"), boiler.AppConfig.GetString("jwt.secret"))
		customerLoginResponseJSON, _ := json.Marshal(customerLoginResponse)
		boiler.Cache.Del(constants.CACHE_CUSTOMER_LOGIN + existingCustomer.Email.String)
		err := boiler.Cache.SetNX(constants.CACHE_CUSTOMER_LOGIN+existingCustomer.Email.String, customerLoginResponseJSON, boiler.AppConfig.GetDuration("cache.expireactivation")*time.Second).Err()
		if err != nil {
			panic(err)
		}
		customerLoginResponse.UserId = 0
		context.JSON(http.StatusOK, response.BaseResponse{
			Data:    structs.Map(customerLoginResponse),
			Message: constants.SUCCESS_MSG_DATA_FOUND,
			Result:  true,
		})
		return
	}
}
