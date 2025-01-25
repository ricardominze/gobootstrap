package service_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" //Driver SQLite
	"github.com/pressly/goose"
	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
	account_service "github.com/ricardominze/gobootstrap/core/domain/account/service"
	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_service "github.com/ricardominze/gobootstrap/core/domain/customer/service"
	"github.com/ricardominze/gobootstrap/core/valueobject"
	"github.com/ricardominze/gobootstrap/infra/adapter"
	"github.com/stretchr/testify/suite"
)

type AccountServiceIntegrationSuite struct {
	suite.Suite
	db              *sql.DB
	ctx             context.Context
	address         *valueobject.Address
	account         *account_entity.Account
	customer        *customer_entity.Customer
	accountService  *account_service.AccountService
	customerService *customer_service.CustomerService
}

// Criando a Suite de Testes
func TestAccountServiceIntegrationSuite(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		t.Fatalf("Error loading .env.test file: %v", err)
	}
	suite.Run(t, &AccountServiceIntegrationSuite{})
}

// Inicialização de configurações de objetos e recursos para serem utilizados nos testes da Suite
func (st *AccountServiceIntegrationSuite) SetupSuite() {

	var err error
	st.db, err = sql.Open(os.Getenv("DATABASE"), os.Getenv("DBSTRING"))

	if err == nil {
		goose.SetDialect(os.Getenv("GOOSE_DRIVER"))
		if err = goose.Up(st.db, os.Getenv("MIGRATION_DIR")); err != nil {
			st.T().Fatalf("failed to apply migrations: %v", err)
		}
	}

	//Recursos

	st.ctx = context.Background()
	st.customerService = customer_service.NewCustomerService(adapter.NewCustomerRepository(st.db)) //Seviço de Customer
	st.accountService = account_service.NewAccountService(adapter.NewAccountRepository(st.db))     //Seviço de Account

	//Entidades

	st.address = &valueobject.Address{}
	st.address.City = "Cidade"
	st.address.Street = "Rua"
	st.address.Zipcode = "123456789"

	st.customer = &customer_entity.Customer{}
	st.customer.Name = "Ricardo"
	st.customer.Address = st.address

	st.account = &account_entity.Account{}
	st.account.Balance = 0.00
	st.account.Status = 1
	st.account.TypeAccount = "CP"
	st.account.IdCustomer = st.customer.Id
}

// Destruição de recursos no final dos testes da Suite
func (st *AccountServiceIntegrationSuite) TearDownSuite() {
	st.customerService = nil
	st.accountService = nil
	st.db.Close()
}

// Antes de rodar o corrente teste da Suite
func (st *AccountServiceIntegrationSuite) BeforeTest(suiteName, testName string) {
}

//===========================================
//Caso de Teste: Serviço de Abertura de Conta
//===========================================

func (st *AccountServiceIntegrationSuite) TestAccountServiceOpen() {

	//Create Customer
	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer.Name, st.customer.Name)

	//Ao criar uma conta do tipo Conta-Corrente o cliente recebe um saldo promocional de 50.00

	st.account.TypeAccount = "CC"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err := st.accountService.Open(st.ctx, st.account)

	st.Nil(err)
	st.Equal(account.Balance, 50.00)

	//Ao criar uma conta do tipo Conta-Poupanca o cliente recebe um saldo promocional de 150.00

	st.account.TypeAccount = "CP"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err = st.accountService.Open(st.ctx, st.account)

	st.Nil(err)
	st.Equal(account.Status, 0)
	st.Equal(account.Balance, 150.00)
}

//============================================
//Caso de Teste: Serviço de Recuperar de Conta
//============================================

func (st *AccountServiceIntegrationSuite) TestAccountServiceGet() {

	//Create Customer
	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer.Name, st.customer.Name)

	st.account.TypeAccount = "CC"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err := st.accountService.Open(st.ctx, st.account)
	st.Nil(err)

	account2, _ := st.accountService.Get(st.ctx, account.Id)
	st.NotNil(account2)
}

//===========================================
//Caso de Teste: Serviço de Deposito em Conta
//===========================================

func (st *AccountServiceIntegrationSuite) TestAccountServiceDeposit() {

	//Create Customer
	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer.Name, st.customer.Name)

	st.account.TypeAccount = "CC"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err := st.accountService.Open(st.ctx, st.account)
	st.Nil(err)

	err = st.accountService.Deposit(st.ctx, account, 50.00)
	st.Nil(err)

	balance, _ := st.accountService.Balance(st.ctx, account.Id)
	st.Equal(account.Balance, balance)
}

//===============================
//Caso de Teste: Serviço de Taxes
//===============================

func (st *AccountServiceIntegrationSuite) TestAccountServiceTaxes() {

	//Create Customer
	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer.Name, st.customer.Name)

	st.account.TypeAccount = "CC"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err := st.accountService.Open(st.ctx, st.account)
	st.Nil(err)

	err = st.accountService.Deposit(st.ctx, account, 50.00)
	st.Nil(err)

	st.accountService.Taxes(st.ctx, account)
	balance, _ := st.accountService.Balance(st.ctx, account.Id)
	st.Equal(balance, 90.00)
}

//===============================
//Caso de Teste: Serviço de Saque
//===============================

func (st *AccountServiceIntegrationSuite) TestAccountServiceWithdraw() {

	//Create Customer
	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer.Name, st.customer.Name)

	st.account.TypeAccount = "CC"
	st.account.IdCustomer = customer.Id
	st.account.Status = 0
	account, err := st.accountService.Open(st.ctx, st.account)
	st.Nil(err)

	err = st.accountService.Deposit(st.ctx, account, 50.00)
	st.Nil(err)

	err = st.accountService.Withdraw(st.ctx, account, 30.00)

	balance, _ := st.accountService.Balance(st.ctx, account.Id)
	st.Equal(balance, 70.00)
}

//=======================================
//Caso de Teste: Serviço de Transferencia
//=======================================

func (st *AccountServiceIntegrationSuite) TestAccountServiceTransfer() {

	//Create Customer
	customer1, err := st.customerService.Create(st.ctx, &customer_entity.Customer{Name: "Customer A", Address: &valueobject.Address{City: "City A", Street: "Street A", Zipcode: "000001"}})
	st.Nil(err)

	account1, err := st.accountService.Open(st.ctx, &account_entity.Account{IdCustomer: customer1.Id, TypeAccount: "CC"})
	st.Nil(err)

	customer2, err := st.customerService.Create(st.ctx, &customer_entity.Customer{Name: "Customer B", Address: &valueobject.Address{City: "City B", Street: "Street B", Zipcode: "000002"}})
	st.Nil(err)

	account2, err := st.accountService.Open(st.ctx, &account_entity.Account{IdCustomer: customer2.Id, TypeAccount: "CC"})
	st.Nil(err)

	err = st.accountService.Transfer(st.ctx, account1, account2, 10.00)
	st.Nil(err)

	balance1, err := st.accountService.Balance(st.ctx, account1.Id)
	st.Nil(err)
	st.Equal(balance1, 40.00)

	balance2, err := st.accountService.Balance(st.ctx, account2.Id)
	st.Nil(err)
	st.Equal(balance2, 60.00)
}
