//go:build wireinjection
// +build wireinjection

package account

import (
	"database/sql"

	"github.com/google/wire"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	account_service "github.com/ricardominze/gobootstrap/core/domain/account/service"
	"github.com/ricardominze/gobootstrap/infra/adapter"
)

func NewAccountDependenciesInjection(db *sql.DB) *account_service.AccountService {
	wire.Build(
		adapter.NewAccountRepository,
		wire.Bind(new(account_port.AccountIRepository), new(*adapter.AccountRepository)),
		account_service.NewAccountService,
	)
	return &account_service.AccountService{}
}
