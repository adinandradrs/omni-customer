package repository

import (
	cbase "github.com/adinandradrs/codefun-go-service"
	"github.com/adinandradrs/omni-customer/model/entity"
)

type customerCapsule struct {
	super cbase.BaseRepositoryCapsule
}

type CustomerRepository interface {
	Register(ncust entity.Customer) error
	Activate(phone string) error
	FindByPhone(phone string) (out entity.Customer, err error)
}

func NewCustomerRepository(super cbase.BaseRepositoryCapsule) CustomerRepository {
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
