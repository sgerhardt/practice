package time

import (
	"fmt"
	"time"
)

func daysInMonth(t time.Time) int {
	currentYear, currentMonth, _ := t.Date()
	currentLocation := t.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	return lastOfMonth.Day()
}
