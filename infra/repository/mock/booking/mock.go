package booking

import (
	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type Mock struct {
	bookings []*restaurant.Booking
}

func New() *Mock {
	return &Mock{
		bookings: make([]*restaurant.Booking, 0),
	}
}

func (m *Mock) Create(booking *restaurant.Booking) error {
	m.bookings = append(m.bookings, booking)
	return nil
}
