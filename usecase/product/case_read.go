package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (s *Service) Read(ctx context.Context, params map[string]string) ([]*entity.Product, error) {
	products, err := s.Repository.Read(ctx, params)
	if err != nil {
		s.Logger.Error(ctx, err.Error())
		return nil, ErrProductListError
	}

	return products, nil
}