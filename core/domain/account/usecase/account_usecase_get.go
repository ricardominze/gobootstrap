package account_usecase

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountUseCaseGet struct {
	repository account_port.AccountIRepository
}

func NewAccountUseCaseGet(repository account_port.AccountIRepository) *AccountUseCaseGet {
	return &AccountUseCaseGet{repository: repository}
}

func (o *AccountUseCaseGet) Execute(ctx context.Context, id int) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	account, err := o.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
