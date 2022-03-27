package restaurant

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/usecase/repository"
)

type usecaseCreateBooking struct {
	bookingRepository repository.Booking
}

type CreateBookingInput struct {
	Username     string `json:"username"`
	CustomerName string `json:"customer_name"`
	BookingDate  string `json:"booking_date"`
	TableID      int    `json:"table_id"`
}

func NewCreateBooking(bookingRepository repository.Booking) *usecaseCreateBooking {
	return &usecaseCreateBooking{
		bookingRepository: bookingRepository,
	}
}

func (u *usecaseCreateBooking) Perform(ctx context.Context, input CreateBookingInput) error {
	bookingDate, err := time.Parse(time.RFC3339, input.BookingDate)
	if err != nil {
		return errors.New("invalid date")
	}

	if bookingDate.Before(time.Now()) {
		return errors.New("can't create a booking at a past date")
	}

	bookings, err := u.bookingRepository.GetBookingsByDay(ctx, bookingDate)
	if err != nil {
		return err
	}

	isAvailable := false
	if len(bookings) == 0 {
		isAvailable = true
	}

	for _, booking := range bookings {
		if booking.TableID == input.TableID &&
			bookingDate.After(booking.Date.Add(2*time.Hour)) {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return errors.New("table not available for booking")
	}

	bookingTime := fmt.Sprintf("%d:%d", bookingDate.Hour(), bookingDate.Minute())

	isWorkingHour := false
	switch bookingDate.Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday:
		isWorkingHour = bookingTime >= "11:00" && bookingTime <= "15:00"
	case time.Thursday, time.Friday, time.Saturday, time.Sunday:
		isWorkingHour = bookingTime >= "11:00" && bookingTime <= "22:00"
	}

	if !isWorkingHour {
		return errors.New("not working datetime")
	}

	isBookingHour := false
	switch bookingDate.Weekday() {
	case time.Thursday, time.Friday, time.Saturday, time.Sunday:
		isBookingHour = bookingTime >= "11:00" && bookingTime <= "20:00"
	}

	if !isBookingHour {
		return errors.New("not booking datetime")
	}

	err = u.bookingRepository.Create(ctx, &restaurant.Booking{
		Username:     input.Username,
		CustomerName: input.CustomerName,
		Date:         bookingDate,
		TableID:      input.TableID,
	})
	if err != nil {
		return err
	}

	return nil
}
