package restaurant

import "time"

type Booking struct {
	Username     string
	CustomerName string
	Date         time.Time
	TableID      int
}
