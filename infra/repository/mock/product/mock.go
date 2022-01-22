package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

type Mock struct {
	products []*entity.Product

	onCreate func(ctx context.Context, product *entity.Product) error
	onRead   func(ctx context.Context, params map[string]string) ([]*entity.Product, error)
	onUpdate func(ctx context.Context, product *entity.Product) error
	onDelete func(ctx context.Context, product *entity.Product) error
}

func (m *Mock) Create(ctx context.Context, product *entity.Product) error {
	if m.onCreate != nil {
		return m.onCreate(ctx, product)
	}
	return nil
}

func (m *Mock) Read(ctx context.Context, params map[string]string) ([]*entity.Product, error) {
	if m.onRead != nil {
		return m.onRead(ctx, params)
	}
	return nil, nil
}

func (m *Mock) Update(ctx context.Context, product *entity.Product) error {
	if m.onUpdate != nil {
		return m.onUpdate(ctx, product)
	}
	return nil
}

func (m *Mock) Delete(ctx context.Context, product *entity.Product) error {
	if m.onDelete != nil {
		return m.onDelete(ctx, product)
	}
	return nil
}
