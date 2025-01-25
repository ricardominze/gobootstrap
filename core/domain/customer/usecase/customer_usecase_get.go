package customer_usecase

import (
	"context"

	entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
	customer_port "github.com/ricardominze/gobootstrap/core/domain/customer/port"
	"github.com/ricardominze/gobootstrap/infra/telemetry"
)

type CustomerUseCaseGet struct {
	repository customer_port.CustomerIRepository
}

func NewCustomerUseCaseGet(repository customer_port.CustomerIRepository) *CustomerUseCaseGet {
	return &CustomerUseCaseGet{repository: repository}
}

func (o *CustomerUseCaseGet) Execute(ctx context.Context, id int) (*entity.Customer, error) {

	ctx, span := telemetry.MakeTraceCall(ctx)
	defer span.End()

	customer, err := o.repository.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
