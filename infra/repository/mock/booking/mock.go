package booking

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type Mock struct {
	Bookings []*restaurant.Booking

	GetBookingsFromDayFunc func(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}

func New() *Mock {
	return &Mock{
		Bookings: make([]*restaurant.Booking, 0),
	}
}

func (m *Mock) Create(ctx context.Context, booking *restaurant.Booking) error {
	booking.ID = strconv.Itoa(len(m.Bookings) + 1)
	m.Bookings = append(m.Bookings, booking)
	return nil
}

func (m *Mock) GetBookingsForDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	results := make([]*restaurant.Booking, 0)
	for i, booking := range m.Bookings {
		if booking.Status != "deleted" &&
			booking.Date.Year() == bookingDate.Year() &&
			booking.Date.Month() == bookingDate.Month() &&
			booking.Date.Day() == bookingDate.Day() {
			results = append(results, m.Bookings[i])
		}
	}

	return results, nil
}

func (m *Mock) GetBookingsFromDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error) {
	if m.GetBookingsFromDayFunc != nil {
		return m.GetBookingsFromDayFunc(ctx, bookingDate)
	}

	results := make([]*restaurant.Booking, 0)
	for i, booking := range m.Bookings {
		if booking.Status != "deleted" &&
			booking.Date.Year() == bookingDate.Year() &&
			booking.Date.Month() == bookingDate.Month() &&
			booking.Date.Day() >= bookingDate.Day() {
			results = append(results, m.Bookings[i])
		}
	}

	return results, nil
}

func (m *Mock) Delete(ctx context.Context, id string) error {
	for i, booking := range m.Bookings {
		if booking.ID == id {
			m.Bookings[i].Status = "deleted"
			return nil
		}
	}

	return errors.New("booking not found")
}
