package helper

import (
	"errors"
	"time"
)

const (
	// YYYYMMDD is the format used to parse dates in the API
	// Example:
	// 		date := "2022-04-16"
	YYYYMMDD = "2006-01-02"
)

// RemoveTimeFromDate removes the time from the date
// Example:
// 		date := time.Date(2022, time.April, 16, 12, 00, 00, 0, time.UTC)
// 		dateWithoutTime := RemoveTimeFromDate(date)
// 		dateWithoutTime.String() == "2022-04-16 00:00:00 +0000 UTC"
func RemoveTimeFromDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// SetDateToLastMomentOfTheDay sets the date to the last moment of the day
// Example:
// 		date := time.Date(2022, time.April, 16, 12, 00, 00, 0, time.UTC)
// 		dateAtLastMoment := SetDateToLastMomentOfTheDay(date)
// 		dateAtLastMoment.String() == "2022-04-16 23:59:59 +0000 UTC"
func SetDateToLastMomentOfTheDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())
}

// TryParseDateToFormats tries to parse the date to the given formats
// Example:
// 		date := "2022-04-16"
// 		dateFormatted, _ := TryParseDateToFormats(date, time.RFC3339, helper.YYYYMMDD)
// 		dateFormatted.String() == "2022-04-16 00:00:00 +0000 UTC"
// Example:
// 		date := "2022-04-16T12:00:00-03:00"
// 		dateFormatted, _ := TryParseDateToFormats(date, helper.YYYYMMDD)
// 		dateFormatted.String() == "2022-04-16"
func TryParseDateToFormats(date string, formats ...string) (time.Time, error) {
	for _, format := range formats {
		t, err := time.Parse(format, date)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("cannot parse date to any of the formats")
}
