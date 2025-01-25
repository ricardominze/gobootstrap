//go:build wireinjection
// +build wireinjection

package customer

import (
	"database/sql"

	"github.com/google/wire"
	customer_port "github.com/ricardominze/gobootstrap/core/domain/customer/port"
	customer_service "github.com/ricardominze/gobootstrap/core/domain/customer/service"
	"github.com/ricardominze/gobootstrap/infra/adapter"
)

func NewCustomerDependenciesInjection(db *sql.DB) *customer_service.CustomerService {
	wire.Build(
		adapter.NewCustomerRepository,
		wire.Bind(new(customer_port.CustomerIRepository), new(*adapter.CustomerRepository)),
		customer_service.NewCustomerService,
	)
	return &customer_service.CustomerService{}
}
