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
	"gopkg.in/inconshreveable/log15.v2"
)

func (repository DatabaseHandler) CustomerActivation(context *gin.Context) {
	var input request.CustomerActivationRequest
	if error := context.ShouldBindJSON(&input); error != nil {
		context.JSON(http.StatusBadRequest, response.BaseResponse{
			Data:    nil,
			Message: error.Error(),
			Result:  false,
		})
		return
	}
	var existingCustomer entity.Customer
	if result := repository.DB.Where("activation_id = ?", input.ActivationId).Where("status = ?", constants.CUSTOMER_STATUS_REGISTERED).First(&existingCustomer); result.Error != nil {
		log15.Error("Customer activation failed to find with error = ", result.Error)
		context.JSON(http.StatusNotFound, response.BaseResponse{
			Data:    nil,
			Message: constants.ERR_MSG_DATA_NOT_FOUND,
			Result:  false,
		})
		return
	}
	existingCustomer.Status = constants.CUSTOMER_STATUS_ACTIVATED
	existingCustomer.ActivationDate = sql.NullTime{Time: time.Now()}
	repository.DB.Save(existingCustomer)
	context.JSON(http.StatusInternalServerError, response.BaseResponse{
		Data:    nil,
		Message: constants.SUCCESS_MSG_DATA_SUBMIT,
		Result:  true,
	})
}
