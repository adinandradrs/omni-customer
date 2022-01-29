package controller

import (
	"net/http"

	cbase "github.com/adinandradrs/codefun-go-common"
	"github.com/adinandradrs/codefun-go-service/util"
	"github.com/adinandradrs/omni-customer/model"
	"github.com/adinandradrs/omni-customer/service"
	"github.com/gin-gonic/gin"
)

type customerControllerCapsule struct {
	activation service.CustomerActivation
	register   service.CustomerRegister
}

type CustomerController interface {
	Activate(ctx *gin.Context)
	Register(c *gin.Context)
}

func NewCustomerController(activation service.CustomerActivation, register service.CustomerRegister) CustomerController {
	return &customerControllerCapsule{
		activation: activation,
		register:   register,
	}
}

func (c *customerControllerCapsule) Activate(ctx *gin.Context) {
	var i model.CustomerActivationRequest
	if err := ctx.ShouldBindJSON(&i); err != nil {
		util.ThrowBadError(err.Error(), ctx)
		return
	}
	resp, err := c.activation.Execute(i)
	if err != nil {
		util.ThrowAnyError(resp, ctx)
		return
	}
	ctx.JSON(http.StatusOK, cbase.RestResponse{
		Message: util.SUCCESS_MSG_DATA_SUBMIT,
		Result:  true,
		Data:    resp,
	})
}

func (c *customerControllerCapsule) Register(ctx *gin.Context) {
	var i model.CustomerRegisterRequest
	if err := ctx.ShouldBindJSON(&i); err != nil {
		util.ThrowBadError(err.Error(), ctx)
		return
	}
	resp, err := c.activation.Execute(i)
	if err != nil {
		util.ThrowAnyError(resp, ctx)
		return
	}
	ctx.JSON(http.StatusOK, cbase.RestResponse{
		Message: util.SUCCESS_MSG_DATA_SUBMIT,
		Result:  true,
		Data:    resp,
	})
}
