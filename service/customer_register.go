package service

import (
	cbase "github.com/adinandradrs/boiler-go-common"
	"github.com/adinandradrs/omni-customer/model"
	"github.com/adinandradrs/omni-customer/model/entity"
	"github.com/adinandradrs/omni-customer/repository"
	"github.com/mitchellh/mapstructure"
)

type customerRegisterCapsule struct {
	custrepo repository.CustomerRepository
}

type CustomerRegister interface {
	cbase.BaseService
}

func NewCustomerRegister(custrepo repository.CustomerRepository) CustomerRegister {
	return &customerRegisterCapsule{
		custrepo: custrepo,
	}
}

func (c *customerRegisterCapsule) Execute(input interface{}) (interface{}, error) {
	req := model.CustomerRegisterRequest{}
	mapstructure.Decode(input, &req)
	c.custrepo.Register(entity.Customer{})
	return model.CustomerRegisterResponse{}, nil
}
