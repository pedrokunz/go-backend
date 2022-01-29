package product

import "github.com/pedrokunz/go_backend/usecase"

type Service struct {
	Repository Repository
	Logger     usecase.Logger
}

func NewService(repository Repository, logger usecase.Logger) (*Service, error) {
	if repository == nil {
		return nil, usecase.ErrRepositoryIsNil
	}

	if logger == nil {
		return nil, usecase.ErrLoggerIsNil
	}
	
	return &Service{
		Logger:     logger,
		Repository: repository,
	}, nil
}