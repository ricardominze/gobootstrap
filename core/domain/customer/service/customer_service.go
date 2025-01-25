package customer_service

import (
	"context"

	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_port "github.com/ricardominze/gobootstrap/core/domain/customer/port"
	customer_usecase "github.com/ricardominze/gobootstrap/core/domain/customer/usecase"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type CustomerService struct {
	ucCustomerGet    *customer_usecase.CustomerUseCaseGet
	ucCustomerCreate *customer_usecase.CustomerUseCaseCreate
}

func NewCustomerService(repository customer_port.CustomerIRepository) *CustomerService {
	return &CustomerService{
		ucCustomerGet:    customer_usecase.NewCustomerUseCaseGet(repository),
		ucCustomerCreate: customer_usecase.NewCustomerUseCaseCreate(repository),
	}
}

func (o *CustomerService) Get(ctx context.Context, id int) (*customer_entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	return o.ucCustomerGet.Execute(ctx, id)
}

func (o *CustomerService) Create(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	customer, err := o.ucCustomerCreate.Execute(ctx, customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (o *CustomerService) Save(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	customer, err := o.ucCustomerCreate.Execute(ctx, customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
