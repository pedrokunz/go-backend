package repository

import (
	"context"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type CreateBooking interface {
	Create(ctx context.Context, input *restaurant.Booking) error

	GetBookingsForDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}

type ListBooking interface {
	GetBookingsFromDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}

type DeleteBooking interface {
	Delete(ctx context.Context, id string) error
}