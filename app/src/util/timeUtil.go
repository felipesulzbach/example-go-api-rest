package util

import (
	"log"
	"time"
)

// FormatDateTime - Returns the formatted date 'yyyy/mm/dd hh:mm:ss'.
func FormatDateTime(dateTime time.Time) string {
	// Visit the 'https://golang.org/src/time/format.go' link for more formats.
	dateFormated := dateTime.Format("2006/01/02 15:04:05")
	return dateFormated
}

// StringToTime - Converts date from String to Time.
func StringToTime(date string) time.Time {
	datetime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Panic(err)
	}
	return datetime
}
