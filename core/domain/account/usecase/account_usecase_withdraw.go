package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseWithdraw struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseWithdraw(repository account_port.AccountIRepository) *AccountUseCaseWithdraw {
	return &AccountUseCaseWithdraw{repository: repository}
}

func (o *AccountUseCaseWithdraw) Execute(ctx context.Context, account *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := account.IsValid()

	if err != nil {
		return err
	}

	//Sacar valor da conta.
	err = account.Withdraw(value)

	if err != nil {
		return err
	}

	_, err = o.repository.Save(ctx, account)

	if err != nil {
		return err
	}

	return nil
}
