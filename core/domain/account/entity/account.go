package account_entity

import (
	errord "github.com/ricardominze/gobootstrap/core/domain/account/err"
)

type Account struct {
	Id          int
	IdCustomer  int
	TypeAccount string
	Balance     float64
	Status      int
}

func NewAccount(id int, typeAccount string, idCustomer int) *Account {

	return &Account{
		Id:          id,
		IdCustomer:  idCustomer,
		TypeAccount: typeAccount,
	}
}

func (o *Account) IsValid() error {

	return nil
}

func (o *Account) Taxes() error {

	var value float64

	if o.TypeAccount == "CC" {
		value = 10.00
	}

	if o.TypeAccount == "CP" {
		value = 12.00
	}

	if o.Balance < value {
		return errord.AccountErrorInsufficientBalance
	}
	o.Balance -= value
	return nil
}

func (o *Account) Deposit(value float64) error {
	if o.Status == 1 {
		return errord.AccountErrorDepositClosed
	}
	o.Balance += value
	return nil
}

func (o *Account) Withdraw(value float64) error {
	if o.Balance < value {
		return errord.AccountErrorInsufficientBalance
	}
	o.Balance -= value
	return nil
}

func (o *Account) CloseAccount() error {

	if o.Balance > 0.0 {
		return errord.AccountErrorClosePositive
	}

	if o.Balance < 0.0 {
		return errord.AccountErrorCloseNegative
	}

	o.Status = 1

	return nil
}
