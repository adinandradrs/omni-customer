package repository

import (
	"github.com/adinandradrs/codefun-go-service/base"
	"github.com/adinandradrs/omni-customer/model/entity"
)

type customerCapsule struct {
	super base.BaseRepositoryCapsule
}

type CustomerRepository interface {
	Register(ncust entity.Customer) error
	Activate(phone string) error
	FindByPhone(phone string) (out entity.Customer, err error)
}

func NewCustomerRepository(super base.BaseRepositoryCapsule) CustomerRepository {
	return &customerCapsule{
		super: super,
	}
}

func (c *customerCapsule) Register(ncust entity.Customer) error {
	return nil
}

func (c *customerCapsule) Activate(phone string) error {
	return nil
}

func (c *customerCapsule) FindByPhone(phone string) (out entity.Customer, err error) {
	extcustomer := entity.Customer{}
	return extcustomer, nil
}
