package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseClose struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseClose(repository account_port.AccountIRepository) *AccountUseCaseClose {
	return &AccountUseCaseClose{repository: repository}
}

func (o *AccountUseCaseClose) Execute(ctx context.Context, account *account_entity.Account) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := account.IsValid()

	if err != nil {
		return err
	}

	//Fechar conta.
	err = account.CloseAccount()

	if err != nil {
		return err
	}

	_, err = o.repository.Save(ctx, account)

	if err != nil {
		return err
	}

	return nil
}
