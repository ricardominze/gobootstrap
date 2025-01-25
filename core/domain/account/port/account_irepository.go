package account_port

import (
	"context"

	account_entity "github.com/ricardominze/gobootstrap/core/domain/account/entity"
)

type AccountIRepository interface {
	Get(ctx context.Context, id int) (*account_entity.Account, error)
	Save(ctx context.Context, account *account_entity.Account) (*account_entity.Account, error)
}
