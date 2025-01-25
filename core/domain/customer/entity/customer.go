package customer_entity

import (
	errord "github.com/ricardominze/gobootstrap/core/domain/customer/err"
	"github.com/ricardominze/gobootstrap/core/valueobject"
)

type Customer struct {
	Id      int
	Name    string
	Address *valueobject.Address
}

func NewCustomer(id int, name string, address *valueobject.Address) *Customer {

	return &Customer{
		Id:      id,
		Name:    name,
		Address: address,
	}
}

func (o *Customer) IsValid() error {

	if len(o.Name) == 0 {
		return errord.CustomerErrorEmptyName
	}

	return nil
}

func (o *Customer) ChangeName(name string) {

	o.Name = name
}
