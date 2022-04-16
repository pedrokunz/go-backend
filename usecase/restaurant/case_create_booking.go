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
	bookingRepository repository.CreateBooking
}

type CreateBookingInput struct {
	Username     string `json:"username"`
	CustomerName string `json:"customer_name"`
	BookingDate  string `json:"booking_date"`
	TableID      int    `json:"table_id"`
}

func NewCreateBooking(bookingRepository repository.CreateBooking) *usecaseCreateBooking {
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

	isAvailable := restaurant.BookingIsAvailable(bookings, input.TableID, bookingDate)
	if !isAvailable {
		return errors.New("table not available for booking")
	}

	bookingTime := fmt.Sprintf("%d:%d", bookingDate.Hour(), bookingDate.Minute())

	err = u.checkHourAvailability(bookingDate, bookingTime)
	if err != nil {
		return err
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

func (u *usecaseCreateBooking) checkHourAvailability(bookingDate time.Time, bookingTime string) error {
	isBookingHour := false
	isWorkingHour := false
	beginShiftTime := "11:00"
	endShiftTime1 := "15:00"
	endShiftTime2 := "22:00"
	endBookingTime := "20:00"

	switch bookingDate.Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday:
		isWorkingHour = bookingTime >= beginShiftTime && bookingTime <= endShiftTime1
	case time.Thursday, time.Friday, time.Saturday, time.Sunday:
		isWorkingHour = bookingTime >= beginShiftTime && bookingTime <= endShiftTime2
		isBookingHour = bookingTime >= beginShiftTime && bookingTime <= endBookingTime
	}

	if !isWorkingHour {
		return errors.New("not working datetime")
	}

	if !isBookingHour {
		return errors.New("not booking datetime")
	}

	return nil
}
