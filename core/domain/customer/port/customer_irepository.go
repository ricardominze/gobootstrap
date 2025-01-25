package customer_port

import (
	"context"

	customer_entity "github.com/ricardominze/gobootstrap/core/domain/customer/entity"
)

type CustomerIRepository interface {
	Get(ctx context.Context, id int) (*customer_entity.Customer, error)
	Save(ctx context.Context, customer *customer_entity.Customer) (*customer_entity.Customer, error)
}
