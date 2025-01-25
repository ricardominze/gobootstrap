package account_entity_test

import (
	"testing"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	"github.com/stretchr/testify/suite"
)

type AccountUnitSuite struct {
	suite.Suite
	account *account_entity.Account
}

// =========================
// Criando a Suite de Testes
// =========================
func TestAccountUnitSuite(t *testing.T) {
	suite.Run(t, &AccountUnitSuite{})
}

// Inicialização de configurações de objetos e recursos para serem utilizados nos testes da Suite
func (st *AccountUnitSuite) SetupSuite() {
	st.account = &account_entity.Account{}
}

// Destruição de recursos no final dos testes da Suite
func (st *AccountUnitSuite) TearDownSuite() {
	st.account = nil
}

// Antes de rodar o corrente teste da Suite
func (st *AccountUnitSuite) BeforeTest(suiteName, testName string) {
}

//==================================
//Caso de Teste: Pagamento de Tarifa
//==================================

type payTest struct {
	balance float64
	message string
}

func (st *AccountUnitSuite) TestAccountEntityTaxes() {

	account := st.account

	var addPays = []payTest{
		{0.00, "Desconto efetuado em conta com saldo 0.00"},
		{10.00, "Desconto efetuado em conta com saldo insuficiente"},
		{20.00, "Desconto não efetuado em conta"},
	}

	for _, test := range addPays {
		account.Balance = test.balance
		balanceBeforePay := account.Balance
		account.Taxes()
		st.GreaterOrEqualf(balanceBeforePay, account.Balance, "==> %s <==", test.message)
	}
}

type payTypeAccountTest struct {
	balance     float64
	tariff      float64
	message     string
	typeAccount string
}

func (st *AccountUnitSuite) TTestAccountEntityIsValid() {

	account := st.account
	st.Nil(account.IsValid())
}

func (st *AccountUnitSuite) TestAccountEntityTaxesType() {

	account := st.account

	var addPaysTypeAccount = []payTypeAccountTest{
		{20.00, 10.00, "Valor de desconto em C/C errado", "CC"},
		{20.00, 12.00, "Valor de desconto em C/P errado", "CP"},
		{10.00, 12.00, "Teste", "CP"},
	}

	for _, test := range addPaysTypeAccount {
		account.TypeAccount = test.typeAccount
		account.Balance = test.balance
		balanceBeforePay := account.Balance
		account.Taxes()

		if test.balance > test.tariff {
			st.Equal((balanceBeforePay - account.Balance), test.tariff, "==> %s <==", test.message)
		} else {
			st.GreaterOrEqual(balanceBeforePay, account.Balance, test.tariff, "==> %s <==", test.message)
		}
	}
}

//==================================
//Caso de Teste: Fechamento de Conta
//==================================

type closeTest struct {
	balance float64
	message string
}

func (st *AccountUnitSuite) TestAccountEntityClose() {

	var err error
	account := st.account

	var addCloses = []closeTest{
		{-10.00, "Conta fechada com saldo negativo"},
		{10.00, "Saque efetuado sem saldo positivo"},
		{0.00, "Conta sem saldo mas, não foi possivel fecha-la"},
	}

	for _, test := range addCloses {
		account.Balance = test.balance
		err = account.CloseAccount()

		if test.balance != 0.00 {
			st.NotNilf(err, "==> %s <==", test.message)
		} else {
			st.Nilf(err, "==> %s <==", test.message)
		}
	}
}

//=======================
//Caso de Teste: Deposito
//=======================

type depositTest struct {
	status  int
	balance float64
	deposit float64
	message string
}

func (st *AccountUnitSuite) TestAccountEntityDeposit() {

	var err error
	account := account_entity.Account{}

	var addDeposits = []depositTest{
		{1, 0.00, 10.00, "Deposito efetuado em conta encerrada"},
		{1, -10.00, 10.00, "Deposito não realizado em conta aberta"},
		{0, 10.00, 10.00, "Deposito não realizado em conta aberta"},
	}

	for _, test := range addDeposits {
		account.Status = test.status
		account.Balance = test.balance
		err = account.Deposit(test.deposit)
		if test.status == 1 {
			st.NotNilf(err, "==> %s <==", test.message)
		} else {
			st.Nilf(err, "==> %s <==", test.message)
		}
	}
}

//====================
//Caso de Teste: Saque
//====================

type withdrawTest struct {
	balance  float64
	withdraw float64
	message  string
}

func (st *AccountUnitSuite) TestAccountEntityWithdraw() {

	var err error
	account := account_entity.Account{}

	var addWithdraws = []withdrawTest{
		{0.00, 10.00, "Saque efetuado sem saldo disponível"},
		{-10.00, 10.00, "Saque efetuado sem saldo negativo"},
		{10.00, 20.00, "Saque efetuado com saldo inferior ao solicitado"},
		{30.00, 20.00, "Saque efetuado com saldo inferior ao solicitado"},
	}

	for _, test := range addWithdraws {
		account.Balance = test.balance
		err = account.Withdraw(test.withdraw)

		if test.balance <= test.withdraw {
			st.NotNilf(err, "==> %s <==", test.message)
		} else {
			st.Nilf(err, "==> %s <==", test.message)
		}
	}
}
