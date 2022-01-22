package product

import "github.com/pedrokunz/go_backend/usecase"

type Service struct {
	Repository Repository
	Logger     usecase.Logger
}

func NewService(repository Repository, logger usecase.Logger) *Service {
	return &Service{
		Logger:     logger,
		Repository: repository,
	}
}