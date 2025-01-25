package account_usecase

import (
	"context"

	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseBalance struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseBalance(repository account_port.AccountIRepository) *AccountUseCaseBalance {
	return &AccountUseCaseBalance{repository: repository}
}

func (o *AccountUseCaseBalance) Execute(ctx context.Context, idAccount int) (float64, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	account, err := o.repository.Get(ctx, idAccount)

	if err != nil {
		return 0.00, err
	}

	return account.Balance, nil
}
