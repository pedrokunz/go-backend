package restaurant

import "time"

type Booking struct {
	ID           string
	Username     string
	CustomerName string
	Date         time.Time
	TableID      int
	Status       string
}

func BookingIsAvailable(bookings []*Booking, tableID int, date time.Time) bool {
	if len(bookings) == 0 {
		return true
	}

	isAvailable := false
	for _, booking := range bookings {
		if booking.TableID == tableID &&
			(date.Equal(booking.Date.Add(2*time.Hour)) ||
				date.Equal(booking.Date.Add(-2*time.Hour)) ||
				date.After(booking.Date.Add(2*time.Hour)) ||
				date.Before(booking.Date.Add(-2*time.Hour))) {
			isAvailable = true
		} else if booking.TableID != tableID {
			isAvailable = true
		}
	}

	return isAvailable
}
