package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseDeposit struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseDeposit(repository account_port.AccountIRepository) *AccountUseCaseDeposit {
	return &AccountUseCaseDeposit{repository: repository}
}

func (o *AccountUseCaseDeposit) Execute(ctx context.Context, account *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := account.IsValid()

	if err != nil {
		return err
	}

	//Depositar valor na conta.
	err = account.Deposit(value)

	if err != nil {
		return err
	}

	_, err = o.repository.Save(ctx, account)

	if err != nil {
		return err
	}

	return nil
}
