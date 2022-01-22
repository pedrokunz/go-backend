package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
	"github.com/pedrokunz/go_backend/usecase/product"
)

type Mock struct {
	products []*entity.Product

	onCreate func(ctx context.Context, product *entity.Product) error
	onRead   func(ctx context.Context, params map[string]string) ([]*entity.Product, error)
	onUpdate func(ctx context.Context, product *entity.Product) error
	onDelete func(ctx context.Context, product *entity.Product) error
}

func New() product.Repository {
	mock := &Mock{products: make([]*entity.Product, 0)}

	mock.Setup()

	return mock
}

func (mock *Mock) Setup() {
	mock.onCreate = func(_ context.Context, product *entity.Product) error {
		mock.products = append(mock.products, product)
		return nil
	}

	mock.onRead = func(_ context.Context, _ map[string]string) ([]*entity.Product, error) {
		return mock.products, nil
	}

	mock.onUpdate = func(_ context.Context, product *entity.Product) error {
		for i, p := range mock.products {
			if p.ID == product.ID {
				mock.products[i] = product
			}
		}

		return nil
	}

	mock.onDelete = func(_ context.Context, product *entity.Product) error {
		for i, p := range mock.products {
			if p.ID == product.ID {
				mock.products = append(mock.products[:i], mock.products[i+1:]...)
			}
		}

		return nil
	}
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
