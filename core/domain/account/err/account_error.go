package errord

import "errors"

var (
	AccountErrorClosePositive       = errors.New("não é possível fechar esta conta, pois existe saldo")
	AccountErrorCloseNegative       = errors.New("não é possível fechar esta conta, pois está em débito com a instituição")
	AccountErrorDepositClosed       = errors.New("não é possível depositar valores nesta conta, pois a mesma se encontra fechada")
	AccountErrorInsufficientBalance = errors.New("saldo insuficiente para a operação")
)
