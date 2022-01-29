package service

import (
	"github.com/adinandradrs/codefun-go-common/base"
	commodel "github.com/adinandradrs/codefun-go-common/model"
	"github.com/adinandradrs/omni-customer/model"
	"github.com/adinandradrs/omni-customer/repository"
	"github.com/mitchellh/mapstructure"
)

type customerActivationCapsule struct {
	custrepo repository.CustomerRepository
}

type CustomerActivation interface {
	base.BaseService
}

func NewCustomerActivation(custrepo repository.CustomerRepository) CustomerActivation {
	return &customerActivationCapsule{
		custrepo: custrepo,
	}
}

func (c *customerActivationCapsule) Execute(input interface{}) (interface{}, error) {
	req := model.CustomerActivationRequest{}
	mapstructure.Decode(input, &req)
	c.custrepo.Activate(req.PhoneNo)
	return commodel.ValidationResponse{Result: true}, nil
}
