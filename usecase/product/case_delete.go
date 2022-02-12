package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (u *Delete) Perform(ctx context.Context, product *entity.Product) error {
	err := u.UseCase.Repository.Delete(ctx, product)
	if err != nil {
		u.UseCase.Logger.Error(ctx, err.Error())
		return ErrProductDeleteError
	}

	return nil
}
