package service

import (
	cbase "github.com/adinandradrs/boiler-go-common"
	"github.com/adinandradrs/omni-customer/model"
	"github.com/adinandradrs/omni-customer/repository"
	"github.com/mitchellh/mapstructure"
)

type customerActivationCapsule struct {
	custrepo repository.CustomerRepository
}

type CustomerActivation interface {
	cbase.BaseService
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
	return cbase.ValidationResponse{Result: true}, nil
}
