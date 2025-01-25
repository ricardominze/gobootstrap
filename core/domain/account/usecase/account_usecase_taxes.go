package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseTaxes struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseTaxes(repository account_port.AccountIRepository) *AccountUseCaseTaxes {
	return &AccountUseCaseTaxes{repository: repository}
}

func (o *AccountUseCaseTaxes) Execute(ctx context.Context, account *account_entity.Account) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := account.IsValid()

	if err != nil {
		return err
	}

	//Pagamento de faturas.
	err = account.Taxes()

	if err != nil {
		return err
	}

	_, err = o.repository.Save(ctx, account)

	if err != nil {
		return err
	}

	return nil
}
