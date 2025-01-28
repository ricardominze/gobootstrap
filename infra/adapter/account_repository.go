package adapter

import (
	"context"
	"database/sql"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (re *AccountRepository) Get(ctx context.Context, id int) (*account_entity.Account, error) {

	account := &account_entity.Account{}

	result := re.db.QueryRowContext(ctx, "SELECT id, id_customer, type_account, balance, status FROM account WHERE id = $1", id)
	result.Scan(&account.Id, &account.IdCustomer, &account.TypeAccount, &account.Balance, &account.Status)

	if result.Err() != nil {
		return nil, result.Err()
	}

	return account, nil
}

func (re *AccountRepository) Save(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	var err error
	var newAccount *account_entity.Account

	if account.Id == 0 {
		newAccount, err = re.create(ctx, account)
	} else {
		newAccount, err = re.update(ctx, account)
	}

	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

func (re *AccountRepository) create(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	result := re.db.QueryRowContext(ctx, `INSERT INTO account (id_customer, type_account, balance, status) VALUES ($1, $2, $3, $4) RETURNING id`, account.IdCustomer, account.TypeAccount, account.Balance, account.Status)
	result.Scan(&account.Id)

	return account, nil
}

func (re *AccountRepository) update(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	_, err := re.db.ExecContext(ctx, `UPDATE account SET id_customer=$1, type_account=$2, balance=$3, status=$4 WHERE id=$5`, account.IdCustomer, account.TypeAccount, account.Balance, account.Status, account.Id)

	if err != nil {
		return nil, err
	}

	return account, nil
}
