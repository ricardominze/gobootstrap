package account_service

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_port "github.com/ricardominze/gobootstrap/core/domain/account/port"
	account_usecase "github.com/ricardominze/gobootstrap/core/domain/account/usecase"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountService struct {
	ucAccountGet      *account_usecase.AccountUseCaseGet
	ucAccountOpen     *account_usecase.AccountUseCaseOpen
	ucAccountClose    *account_usecase.AccountUseCaseClose
	ucAccountTaxes    *account_usecase.AccountUseCaseTaxes
	ucAccountDeposit  *account_usecase.AccountUseCaseDeposit
	ucAccountBalance  *account_usecase.AccountUseCaseBalance
	ucAccountWithdraw *account_usecase.AccountUseCaseWithdraw
	ucAccountTransfer *account_usecase.AccountUseCaseTransfer
}

func NewAccountService(repository account_port.AccountIRepository) *AccountService {
	return &AccountService{
		ucAccountGet:      account_usecase.NewAccountUseCaseGet(repository),
		ucAccountTaxes:    account_usecase.NewAccountUseCaseTaxes(repository),
		ucAccountOpen:     account_usecase.NewAccountUseCaseOpen(repository),
		ucAccountClose:    account_usecase.NewAccountUseCaseClose(repository),
		ucAccountDeposit:  account_usecase.NewAccountUseCaseDeposit(repository),
		ucAccountWithdraw: account_usecase.NewAccountUseCaseWithdraw(repository),
		ucAccountBalance:  account_usecase.NewAccountUseCaseBalance(repository),
		ucAccountTransfer: account_usecase.NewAccountUseCaseTransfer(repository),
	}
}

func (o *AccountService) Get(ctx context.Context, id int) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountGet.Execute(ctx, id)
}

func (o *AccountService) Taxes(ctx context.Context, account *account_entity.Account) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountTaxes.Execute(ctx, account)
}

func (o *AccountService) Open(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountOpen.Execute(ctx, account)
}

func (o *AccountService) Close(ctx context.Context, account *account_entity.Account) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountClose.Execute(ctx, account)
}

func (o *AccountService) Deposit(ctx context.Context, account *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountDeposit.Execute(ctx, account, value)
}

func (o *AccountService) Withdraw(ctx context.Context, account *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountWithdraw.Execute(ctx, account, value)
}

func (o *AccountService) Transfer(ctx context.Context, accountSource *account_entity.Account, accountDestiny *account_entity.Account, value float64) error {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountTransfer.Execute(ctx, accountSource, accountDestiny, value)
}

func (o *AccountService) Balance(ctx context.Context, idAccount int) (float64, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucAccountBalance.Execute(ctx, idAccount)
}
