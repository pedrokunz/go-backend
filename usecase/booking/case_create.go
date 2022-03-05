package booking

import (
	"errors"
	"github.com/pedrokunz/go_backend/entity/restaurant"
	"github.com/pedrokunz/go_backend/usecase/repository"
	"time"
)

type usecaseCreateBooking struct {
	tableRepository   repository.Table
	hourRepository    repository.Hour
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
	hourRepository repository.Hour,
	bookingRepository repository.Booking) *usecaseCreateBooking {
	return &usecaseCreateBooking{
		tableRepository:   tableRepository,
		hourRepository:    hourRepository,
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

	workingHours, err := u.hourRepository.GetWorkingHours()
	if err != nil {
		return err
	}

	bookingDate, err := time.Parse("2006-01-02 15:04:05", input.BookingDate)
	if err != nil {
		return errors.New("invalid date")
	}

	isWorkingHour := false
	for _, hour := range workingHours {
		startTime, err := time.Parse("2006-01-02 15:04:05", hour.StartTime)
		if err != nil {
			break
		}

		endTime, err := time.Parse("2006-01-02 15:04:05", hour.EndTime)
		if err != nil {
			break
		}

		if bookingDate.After(startTime) && bookingDate.Before(endTime) {
			isWorkingHour = true
			break
		}
	}

	if !isWorkingHour {
		return errors.New("not working hour")
	}

	bookingHours, err := u.hourRepository.GetBookingHours()
	if err != nil {
		return err
	}

	isBookingHour := false
	for _, hour := range bookingHours {
		startTime, err := time.Parse("2006-01-02 15:04:05", hour.StartTime)
		if err != nil {
			break
		}

		endTime, err := time.Parse("2006-01-02 15:04:05", hour.EndTime)
		if err != nil {
			break
		}

		switch bookingDate.Weekday() {
		case time.Thursday, time.Friday, time.Saturday, time.Sunday:
			isBookingHour = true
		}

		if isBookingHour && bookingDate.After(startTime) && bookingDate.Before(endTime) {
			isBookingHour = true
			break
		}
	}

	if !isBookingHour {
		return errors.New("not booking hour")
	}

	err = u.bookingRepository.Create(&restaurant.Booking{})
	if err != nil {
		return err
	}

	return nil
}
