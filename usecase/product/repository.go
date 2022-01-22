package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

type Repository interface {
	Create(ctx context.Context, product *entity.Product) error
	Read(ctx context.Context, params map[string]string) ([]*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, product *entity.Product) error
}
