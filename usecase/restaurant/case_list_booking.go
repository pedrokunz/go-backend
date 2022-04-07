package restaurant

import (
	"context"
	"errors"
	"time"

	"github.com/pedrokunz/go_backend/helper"
	"github.com/pedrokunz/go_backend/usecase/repository"
)

type usecaseListBooking struct {
	bookingRepository repository.ListBooking
}

type ListBookingInput struct {
	Date string `json:"date"`
}

type ListBookingOutput struct {
	Username     string `json:"username"`
	CustomerName string `json:"customer_name"`
	BookingDate  string `json:"booking_date"`
	TableID      int    `json:"table_id"`
}

func NewListBooking(bookingRepository repository.ListBooking) *usecaseListBooking {
	return &usecaseListBooking{
		bookingRepository: bookingRepository,
	}
}

func (u *usecaseListBooking) Perform(ctx context.Context, input ListBookingInput) ([]ListBookingOutput, error) {
	bookingDate, err := helper.TryParseDateToFormats(input.Date, time.RFC3339, helper.YYYYMMDD)
	if err != nil {
		return nil, errors.New("invalid date")
	}

	bookings, err := u.bookingRepository.GetBookingsByDay(ctx, bookingDate)
	if err != nil {
		return nil, err
	}

	bookingsOutput := make([]ListBookingOutput, len(bookings))
	for i, booking := range bookings {
		bookingsOutput[i] = ListBookingOutput{
			Username:     booking.Username,
			CustomerName: booking.CustomerName,
			BookingDate:  booking.Date.Format(time.RFC3339),
			TableID:      booking.TableID,
		}
	}

	return bookingsOutput, nil
}
