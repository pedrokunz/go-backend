package repository

import (
	"github.com/pedrokunz/go_backend/entity/restaurant"
	"time"
)

type Booking interface {
	Create(booking *restaurant.Booking) error

	GetBookingsByDay(bookingDate time.Time) ([]*restaurant.Booking, error)
}
