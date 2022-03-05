package repository

import "github.com/pedrokunz/go_backend/entity/restaurant"

type Hour interface {
	GetWorkingHours() ([]restaurant.Hour, error)
	GetBookingHours() ([]restaurant.Hour, error)
}
