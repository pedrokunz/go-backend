package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (s *Service) Update(ctx context.Context, product *entity.Product) error {
	err := s.Repository.Update(ctx, product)
	if err != nil {
		s.Logger.Error(ctx, err.Error())
		return ErrProductUpdateError
	}

	return nil
}
