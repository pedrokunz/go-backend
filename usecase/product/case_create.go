package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (s *Service) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	err := s.Repository.Create(ctx, product)
	if err != nil {
		s.Logger.Error(ctx, err.Error())
		return nil, ErrProductCreateError
	}

	return product, nil
}