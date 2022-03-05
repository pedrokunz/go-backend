package repository

import "github.com/pedrokunz/go_backend/entity/restaurant"

type Booking interface {
	Create(booking *restaurant.Booking) error
}
