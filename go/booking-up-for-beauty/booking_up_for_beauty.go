package booking

import (
	"fmt"
	"log"
	"time"
)

const SHORTER_DATE_WITH_TIME = "1/2/2006 15:04:05"
const SHORT_DATE_WITH_TIME = "1/02/2006 15:04:05"
const LONG_DATE_WITH_TIME = "January 2, 2006 15:04:05"
const DATE_TIME_USER_FACING_LONG = "Monday, January 2, 2006 15:04:05"
const DATE_TIME_USER_FACING_LONG_WITH_SPACING = "Monday, January 2, 2006, at 15:04."

func Schedule(date string) time.Time {
	time, err := time.Parse(SHORT_DATE_WITH_TIME, date)

	if err != nil {
		log.Fatal(err)
	}

	return time
}

func HasPassed(date string) bool {
	parsedTime, err := time.Parse(LONG_DATE_WITH_TIME, date)

	if err != nil {
		log.Fatal(err)
	}
	return parsedTime.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	parsedTime, err := time.Parse(DATE_TIME_USER_FACING_LONG, date)

	if err != nil {
		log.Fatal(err)
	}

	hour := parsedTime.Hour()

	if hour >= 12 && hour < 18 {
		return true
	} else {
		return false
	}

}

func Description(date string) string {
	parsedTime, err := time.Parse(SHORTER_DATE_WITH_TIME, date)

	if err != nil {
		log.Fatal(err)
	}

	formattedTime := parsedTime.Format(DATE_TIME_USER_FACING_LONG_WITH_SPACING)
	return fmt.Sprintf("You have an appointment on %s", formattedTime)

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
