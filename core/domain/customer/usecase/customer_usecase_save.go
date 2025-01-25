package customer_usecase

import (
	"context"

	entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_port "github.com/ricardominze/gobootstrap/core/domain/customer/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type CustomerUseCaseSave struct {
	repository customer_port.CustomerIRepository
}

func NewCustomerUseCaseSave(repository customer_port.CustomerIRepository) *CustomerUseCaseSave {
	return &CustomerUseCaseSave{repository: repository}
}

func (o *CustomerUseCaseSave) Execute(ctx context.Context, customer *entity.Customer) (*entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	err := customer.IsValid()

	if err != nil {
		return nil, err
	}

	customer, err = o.repository.Save(ctx, customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
