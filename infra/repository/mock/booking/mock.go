package booking

import (
	"context"
	"time"

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

func (m *Mock) Create(ctx context.Context, booking *restaurant.Booking) error {
	m.bookings = append(m.bookings, booking)
	return nil
}

func (m *Mock) GetBookingsByDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	results := make([]*restaurant.Booking, 0)
	for i, booking := range m.bookings {
		if booking.Date.Year() == bookingDate.Year() &&
			booking.Date.Month() == bookingDate.Month() &&
			booking.Date.Day() == bookingDate.Day() {
			results = append(results, m.bookings[i])
		}
	}

	return results, nil
}
