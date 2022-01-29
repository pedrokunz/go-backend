package main

import (
	"github.com/pedrokunz/go_backend/infra/logger/mock"
	"github.com/pedrokunz/go_backend/infra/repository/mock/product"
	productUseCase "github.com/pedrokunz/go_backend/usecase/product"
)

func main() {
	loggerMock := mock.Logger{}

	productRepositoryMock := product.Mock{}

	productService, err := productUseCase.NewService(&productRepositoryMock, &loggerMock)
	if err != nil || productService == nil {
		panic("unable to create product service")
	}
}