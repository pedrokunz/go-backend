package product

import (
	"net/http"

	"github.com/pedrokunz/go_backend/api/model"
)

func New(productHandler model.ProductHandler) http.Handler {

	return productHandler
}
