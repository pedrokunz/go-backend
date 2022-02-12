package product

import "github.com/pedrokunz/go_backend/usecase"

type UseCase struct {
	Repository Repository
	Logger     usecase.Logger
}

type Create struct {
	UseCase UseCase
}

type Read struct {
	UseCase UseCase
}

type Update struct {
	UseCase UseCase
}

type Delete struct {
	UseCase UseCase
}

func NewUseCase(repository Repository, logger usecase.Logger) (UseCase, error) {
	if repository == nil {
		return UseCase{}, usecase.ErrRepositoryIsNil
	}

	if logger == nil {
		return UseCase{}, usecase.ErrLoggerIsNil
	}

	return UseCase{
		Logger:     logger,
		Repository: repository,
	}, nil
}

func NewCreate(useCase UseCase) Create {
	return Create{
		UseCase: useCase,
	}
}

func NewRead(useCase UseCase) Read {
	return Read{
		UseCase: useCase,
	}
}

func NewUpdate(useCase UseCase) Update {
	return Update{
		UseCase: useCase,
	}
}

func NewDelete(useCase UseCase) Delete {
	return Delete{
		UseCase: useCase,
	}
}
