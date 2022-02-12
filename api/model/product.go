package model

import (
	"net/http"

	"github.com/pedrokunz/go_backend/usecase/product"
)

type ProductHandler struct {
	http.Handler
	useCase *product.UseCase
}

type CreateProductInput struct {
	Name string `json:"name"`
}
