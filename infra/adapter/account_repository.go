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

	result := re.db.QueryRowContext(ctx, "SELECT id, id_customer, type_account, balance, status FROM account WHERE id = ?", id)
	result.Scan(&account.Id, &account.IdCustomer, &account.TypeAccount, &account.Balance, &account.Status)

	if result.Err() != nil {
		return nil, result.Err()
	}

	return account, nil
}

func (re *AccountRepository) Save(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	accountOri := &account_entity.Account{}

	result := re.db.QueryRowContext(ctx, "SELECT id, id_customer, type_account, balance, status FROM account WHERE id = ?", account.Id)
	result.Scan(&accountOri.Id, &accountOri.IdCustomer, &accountOri.TypeAccount, &accountOri.Balance, &accountOri.Status)

	if accountOri.Id == 0 {

		_, err := re.create(ctx, account)

		if err != nil {
			return nil, err
		}

	} else {

		_, err := re.update(ctx, account)

		if err != nil {
			return nil, err
		}
	}

	return account, nil
}

func (re *AccountRepository) create(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	result, err := re.db.ExecContext(ctx, "INSERT INTO account (id_customer, type_account, balance, status) VALUES (?, ?, ?, ?)", account.IdCustomer, account.TypeAccount, account.Balance, account.Status)

	if err != nil {
		return nil, err
	}

	lastInsertId, _ := result.LastInsertId()
	account.Id = int(lastInsertId)

	return account, nil
}

func (re *AccountRepository) update(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error) {

	_, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	_, err := re.db.ExecContext(ctx, "UPDATE account SET id_customer=?, type_account=?, balance=?, status=? WHERE id=?", account.IdCustomer, account.TypeAccount, account.Balance, account.Status, account.Id)

	if err != nil {
		return nil, err
	}

	return account, nil
}
