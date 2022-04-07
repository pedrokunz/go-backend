package repository

import (
	"context"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type CreateBooking interface {
	Create(ctx context.Context, input *restaurant.Booking) error

	GetBookingsByDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}

type ListBooking interface {
	GetBookingsByDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}