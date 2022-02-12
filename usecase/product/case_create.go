package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (u *Create) Perform(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	err := u.UseCase.Repository.Create(ctx, product)
	if err != nil {
		u.UseCase.Logger.Error(ctx, err.Error())
		return nil, ErrProductCreateError
	}

	return product, nil
}
