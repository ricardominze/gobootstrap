package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseTransfer struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseTransfer(repository account_port.AccountIRepository) *AccountUseCaseTransfer {
	return &AccountUseCaseTransfer{repository: repository}
}

func (o *AccountUseCaseTransfer) Execute(ctx context.Context, accountSource *account_entity.Account, accountDestiny *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := accountSource.IsValid()

	if err != nil {
		return err
	}

	err = accountDestiny.IsValid()

	if err != nil {
		return err
	}

	//Saque da Conta Origem
	err = accountSource.Withdraw(value)

	if err != nil {
		return err
	}

	//Deposito na Conta Destino
	err = accountDestiny.Deposit(value)

	if err != nil {
		return err
	}

	_, err = o.repository.Save(ctx, accountSource)

	if err != nil {
		return err
	}
	_, err = o.repository.Save(ctx, accountDestiny)

	if err != nil {
		return err
	}

	return nil
}
