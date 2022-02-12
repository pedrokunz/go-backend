package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (u *Update) Update(ctx context.Context, product *entity.Product) error {
	err := u.UseCase.Repository.Update(ctx, product)
	if err != nil {
		u.UseCase.Logger.Error(ctx, err.Error())
		return ErrProductUpdateError
	}

	return nil
}
