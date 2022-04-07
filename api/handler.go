package api

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pedrokunz/go_backend/api/internal/restaurant"
)

func ListenAndServe() {
	h := initGinEngine()

	r := h.Group("/api/v1/restaurant")

	restaurant.HandleCreateBooking(r)
	restaurant.HandleListBooking(r)

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
