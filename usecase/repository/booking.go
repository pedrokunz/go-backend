package repository

import (
	"context"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
)

type Booking interface {
	Create(ctx context.Context, input *restaurant.Booking) error

	GetBookingsByDay(ctx context.Context, bookingDate time.Time) ([]*restaurant.Booking, error)
}
