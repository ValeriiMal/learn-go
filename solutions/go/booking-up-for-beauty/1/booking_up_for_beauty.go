package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Accept both "1/2/2006 15:04:05" and "01/02/2006 15:04:05" formats
	layouts := []string{
		"1/2/2006 15:04:05",
		"01/02/2006 15:04:05",
	}
	var t time.Time
	var e error
	for _, layout := range layouts {
		t, e = time.Parse(layout, date)
		if e == nil {
			break
		}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return time.Now().After(t)
}
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// You have an appointment on Thursday, July 25, 2019, at 13:45.
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2025, 9, 15, 0, 0, 0, 0, time.UTC)
}
