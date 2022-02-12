package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/api/model"
	"github.com/pedrokunz/go_backend/entity"
	mockProductRepo "github.com/pedrokunz/go_backend/infra/repository/mock/product"
	"github.com/pedrokunz/go_backend/usecase/product"
)

func New(httpHandler http.Handler) http.Handler {
	return httpHandler
}

type handler struct {
	http.Handler
	ProductRepository product.Repository
}

func NewHandler() {
	h := gin.Default()
	repo := mockProductRepo.New()

	h.POST("/product", func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		product := model.CreateProductInput{}
		err = json.Unmarshal(body, &product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		products, _ := repo.Read(c.Request.Context(), map[string]string{"name": product.Name})
		if len(products) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("product %s already exists", product.Name),
			})
			return
		}

		repo.Create(c.Request.Context(), &entity.Product{Name: product.Name})

		c.JSON(200, gin.H{
			"message": "product created",
		})
	})

	h.Run()
}
