package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (u *Update) Perform(ctx context.Context, params map[string]string) ([]*entity.Product, error) {
	products, err := u.UseCase.Repository.Read(ctx, params)
	if err != nil {
		u.UseCase.Logger.Error(ctx, err.Error())
		return nil, ErrProductListError
	}

	return products, nil
}
