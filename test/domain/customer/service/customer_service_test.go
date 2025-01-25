package customer_service_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" //Driver SQLite
	"github.com/pressly/goose"
	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_service "github.com/ricardominze/gobootstrap/core/domain/customer/service"
	"github.com/ricardominze/gobootstrap/core/valueobject"
	"github.com/ricardominze/gobootstrap/infra/adapter"
	"github.com/stretchr/testify/suite"
)

type CustomerServiceIntegrationSuite struct {
	suite.Suite
	db              *sql.DB
	ctx             context.Context
	customer        *customer_entity.Customer
	customerService *customer_service.CustomerService
}

// Criando a Suite de Testes
func TestCustomerServiceIntegrationSuite(t *testing.T) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		t.Fatalf("Error loading .env.test file: %v", err)
	}
	suite.Run(t, &CustomerServiceIntegrationSuite{})
}

// Inicialização de configurações de objetos e recursos para serem utilizados nos testes da Suite
func (st *CustomerServiceIntegrationSuite) SetupSuite() {

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

	//Entidades

	address := &valueobject.Address{}
	address.City = "Cidade"
	address.Street = "Rua"
	address.Zipcode = "123456789"

	st.customer = &customer_entity.Customer{Name: "Ricardo", Address: address}
	st.customer.Name = "Ricardo"
	st.customer.Address = address
}

// Destruição de recursos no final dos testes da Suite
func (st *CustomerServiceIntegrationSuite) TearDownSuite() {
	st.db.Close()
}

// Antes de rodar o corrente teste da Suite
func (st *CustomerServiceIntegrationSuite) BeforeTest(suiteName, testName string) {
}

//=============================================
//Caso de Teste: Serviço de Criacao de Cadastro
//=============================================

func (st *CustomerServiceIntegrationSuite) TestCustomerServiceCreate() {

	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)
	st.Equal(customer, st.customer)
}

//=================================================
//Caso de Teste: Serviço de Atualizacao de Cadastro
//=================================================

func (st *CustomerServiceIntegrationSuite) TestCustomerServiceSave() {

	customer, err := st.customerService.Create(st.ctx, st.customer)
	st.Nil(err)

	name := customer.Name
	customer.Name = "Ricardo Minze"
	st.customerService.Save(st.ctx, customer)
	st.NotEqual(name, customer.Name)
}
