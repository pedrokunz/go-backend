package repository

import (
	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type Table interface {
	GetAvailableTables() ([]restaurant.Table, error)
}
