package product

import (
	"context"

	"github.com/pedrokunz/go_backend/entity"
)

func (s *Service) Delete(ctx context.Context, product *entity.Product) error {
	err := s.Repository.Delete(ctx, product)
	if err != nil {
		s.Logger.Error(ctx, err.Error())
		return ErrProductDeleteError
	}

	return nil
}