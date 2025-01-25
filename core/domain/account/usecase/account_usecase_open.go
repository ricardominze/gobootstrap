package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseOpen struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseOpen(repository account_port.AccountIRepository) *AccountUseCaseOpen {
	return &AccountUseCaseOpen{repository: repository}
}

func (o *AccountUseCaseOpen) Execute(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := account.IsValid()

	if err != nil {
		return nil, err
	}

	//Se o tipo de conta for conta corrente, adiciona R$ 50,00 na conta de bônus.
	if account.TypeAccount == "CC" {
		account.Balance = 50.00
	}

	//Se o tipo de conta for conta poupança, adiciona R$ 150,00 na conta de bônus.
	if account.TypeAccount == "CP" {
		account.Balance = 150.00
	}

	account, err = o.repository.Save(ctx, account)

	if err != nil {
		return nil, err
	}

	return account, nil
}
