package controller

import (
	"net/http"
	"omni-customer/model/entity"
	"omni-customer/model/response"
	constants "omni-customer/utility"
	"strconv"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gopkg.in/inconshreveable/log15.v2"
	"gorm.io/gorm"
)

func (boiler ConfigurationHandler) CustomerGetDetail(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	getCustomerDetail(id, "", true, context, boiler)
}

func getCustomerDetail(id int64, email string, idIncluded bool, context *gin.Context, boiler ConfigurationHandler) {
	var existingCustomer entity.Customer
	var findCriteria *gorm.DB
	var selectColumns *gorm.DB
	if id != 0 {
		findCriteria = boiler.DB.Where("id = ?", id)
	}
	if email != "" {
		findCriteria = boiler.DB.Where("email = ?", email)
	}
	if idIncluded {
		selectColumns = boiler.DB.Select("fullname, email, phone_no, address, id")
	} else {
		selectColumns = boiler.DB.Select("fullname, email, phone_no, address")
	}
	if result := selectColumns.
		Where(findCriteria).
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
		Data: structs.Map(response.CustomerDetail{
			Id:       existingCustomer.Id,
			Email:    existingCustomer.Email.String,
			Fullname: existingCustomer.Fullname.String,
			PhoneNo:  existingCustomer.PhoneNo.String,
			Address:  existingCustomer.Address.String,
		}),
		Message: constants.SUCCESS_MSG_DATA_FOUND,
		Result:  true,
	})
}
