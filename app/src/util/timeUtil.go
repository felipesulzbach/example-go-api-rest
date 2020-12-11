package util

import (
	"log"
	"strings"
	"time"

)

// FormatDateTime - Returns the formatted date 'yyyy/mm/dd hh:mm:ss'.
func FormatDateTime(dateTime time.Time) string {
	// Visit the 'https://golang.org/src/time/format.go' link for more formats.
	dateFormated := dateTime.Format("2006/01/02 15:04:05")
	return dateFormated
}

// FormatDateTimeISO8601 - Returns the formatted date in ISO8601 '2006-01-02T15:04:05Z07:00'.
func FormatDateTimeISO8601(dateTime time.Time) string {
	dateFormated := dateTime.UTC().Format("2006-01-02T15:04:05Z07:00")
	return dateFormated
}

// StringToTime - Converts date from String to Time.
func StringToTime(date string) time.Time {
	if strings.Contains(date, " m=") {
		return removeMonotonicClock(date)
	}

	datetime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		datetime, err = time.Parse(time.UnixDate, date)
		if err != nil {
			datetime, err = time.Parse("2006-01-02 15:04:05 -0700 MST", date)
			if err != nil {
				log.Panic(err)
			}
		}
	}
	return datetime
}

// IsDateTimeEmpty ...
func IsDateTimeEmpty(dateTime time.Time) bool {
	return dateTime.IsZero()
}

func removeMonotonicClock(date string) time.Time {
	datetime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", strings.Split(date, " m=")[0])
	if err != nil {
		log.Panic(err)
	}

	return datetime
}

// IsDateTimeBefore ...
func IsDateTimeBefore(start time.Time, end time.Time) bool {
	return start.Before(end)
}
