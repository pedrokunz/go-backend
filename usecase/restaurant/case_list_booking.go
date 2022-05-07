package restaurant

import (
	"context"
	"time"

	"github.com/pedrokunz/go_backend/usecase/repository"
)

type usecaseListBooking struct {
	bookingRepository repository.ListBooking
}

type ListBookingOutput struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	CustomerName string `json:"customer_name"`
	BookingDate  string `json:"booking_date"`
	TableID      int    `json:"table_id"`
	Status       string `json:"status"`
}

func NewListBooking(bookingRepository repository.ListBooking) *usecaseListBooking {
	return &usecaseListBooking{
		bookingRepository: bookingRepository,
	}
}

func (u *usecaseListBooking) Perform(ctx context.Context) ([]ListBookingOutput, error) {
	bookings, err := u.bookingRepository.GetBookingsFromDay(ctx, time.Now())
	if err != nil {
		return nil, err
	}

	bookingsOutput := make([]ListBookingOutput, 0)
	for _, booking := range bookings {
		bookingsOutput = append(bookingsOutput, ListBookingOutput{
			ID:           booking.ID,
			Username:     booking.Username,
			CustomerName: booking.CustomerName,
			BookingDate:  booking.Date.Format(time.RFC3339),
			TableID:      booking.TableID,
			Status:       booking.Status,
		})
	}

	return bookingsOutput, nil
}
