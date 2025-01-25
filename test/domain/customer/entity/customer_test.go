package customer_entity_test

import (
	"testing"

	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"

	"github.com/stretchr/testify/suite"
)

type CustomerUnitSuite struct {
	suite.Suite
	customer *entity.Customer
}

// Criando a Suite de Testes
func TestCustomerUnitSuite(t *testing.T) {
	suite.Run(t, &CustomerUnitSuite{})
}

// Inicialização de configurações de objetos e recursos para serem utilizados nos testes da Suite
func (st *CustomerUnitSuite) SetupSuite() {
	st.customer = &customer_entity.Customer{}
}

// Destruição de recursos no final dos testes da Suite
func (st *CustomerUnitSuite) TearDownSuite() {
	st.customer = nil
}

// Antes de rodar o corrente teste da Suite
func (st *CustomerUnitSuite) BeforeTest(suiteName, testName string) {
}

//==============================
//Caso de Teste: Mudança de Nome
//==============================

type changeNameTest struct {
	name string
}

func (st *CustomerUnitSuite) TestCustomerCreate() {

	customer := st.customer

	var addChangeNames = []changeNameTest{
		{"Ricardo"},
		{"Mario"},
		{"Joao"},
	}

	for _, test := range addChangeNames {
		customer.ChangeName(test.name)
		st.Equal(customer.Name, test.name, "==> %s <==", "Mudança de nome não efetuada")
	}
}
