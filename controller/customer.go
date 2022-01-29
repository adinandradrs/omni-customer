package controller

import (
	"net/http"

	commodel "github.com/adinandradrs/codefun-go-common/model"
	"github.com/adinandradrs/codefun-go-service/util"
	"github.com/adinandradrs/omni-customer/model"
	"github.com/adinandradrs/omni-customer/service"
	"github.com/gin-gonic/gin"
)

type customerControllerCapsule struct {
	activation service.CustomerActivation
	register   service.CustomerRegister
}

type RoleController interface {
	Activate(ctx *gin.Context)
	Register(c *gin.Context)
}

func NewRoleController(activation service.CustomerActivation, register service.CustomerRegister) RoleController {
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
	ctx.JSON(http.StatusOK, commodel.RestResponse{
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
	ctx.JSON(http.StatusOK, commodel.RestResponse{
		Message: util.SUCCESS_MSG_DATA_SUBMIT,
		Result:  true,
		Data:    resp,
	})
}
