package restaurant

import (
	"errors"
	"fmt"
	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/usecase/repository"
	"time"
)

type usecaseCreateBooking struct {
	tableRepository   repository.Table
	bookingRepository repository.Booking
}

type CreateBookingInput struct {
	Username     string
	CustomerName string
	BookingDate  string
	TableID      int
}

func NewCreate(
	tableRepository repository.Table,
	bookingRepository repository.Booking) *usecaseCreateBooking {
	return &usecaseCreateBooking{
		tableRepository:   tableRepository,
		bookingRepository: bookingRepository,
	}
}

func (u *usecaseCreateBooking) Perform(input CreateBookingInput) error {
	tables, err := u.tableRepository.GetAvailableTables()
	if err != nil {
		return err
	}

	isAvailable := false
	for _, table := range tables {
		if table.ID == input.TableID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return errors.New("table not available")
	}

	bookingDate, err := time.Parse(time.RFC3339, input.BookingDate)
	if err != nil {
		return errors.New("invalid date")
	}

	if bookingDate.Before(time.Now()) {
		return errors.New("can't create a booking at a past date")
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

	err = u.bookingRepository.Create(&restaurant.Booking{
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
