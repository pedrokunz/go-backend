package api

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/api/internal/restaurant"
	"github.com/pedrokunz/go_backend/usecase/repository"
)

type ApiDependencies struct {
	ListBookingRepo   repository.ListBooking
	CreateBookingRepo repository.CreateBooking
	DeleteBookingRepo repository.DeleteBooking
}

func ListenAndServe(dependencies ApiDependencies) {
	h := initGinEngine()

	r := h.Group("/api/v1/restaurant")

	restaurant.HandleCreateBooking(r, dependencies.CreateBookingRepo)
	restaurant.HandleListBooking(r, dependencies.ListBookingRepo)
	restaurant.HandleDeleteBooking(r, dependencies.DeleteBookingRepo)

	h.Run(":" + os.Getenv("HTTP_PORT"))
}

func initGinEngine() *gin.Engine {
	h := gin.Default()

	trustedProxiesEnv := os.Getenv("TRUSTED_PROXIES")
	trustedProxies := []string{"localhost"}
	if trustedProxiesEnv != "" {
		trustedProxies = append(trustedProxies, strings.Split(trustedProxiesEnv, ",")...)
	}

	h.SetTrustedProxies(trustedProxies)

	return h
}
