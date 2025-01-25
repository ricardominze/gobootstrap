// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package customer

import (
	"database/sql"
	"github.com/ricardominze/gobootstrap/core/domain/customer/service"
	"github.com/ricardominze/gobootstrap/infra/adapter"
)

// Injectors from customer_di.go:

func NewCustomerDependenciesInjection(db *sql.DB) *customer_service.CustomerService {
	customerRepository := adapter.NewCustomerRepository(db)
	customerService := customer_service.NewCustomerService(customerRepository)
	return customerService
}
