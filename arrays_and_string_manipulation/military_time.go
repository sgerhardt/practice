package arrays_and_string_manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	segments := strings.Split(s, ":")
	hours := segments[0]
	minutes := segments[1]
	seconds := segments[2][0:2] // strip the AM/PM
	hourVal, err := strconv.Atoi(hours)
	if err != nil {
		fmt.Printf("error converting:%+v\n", err)
	}
	if strings.HasSuffix(s, "PM") {
		if hourVal == 12 {
			// do not adjust time
			return strings.Join([]string{hours, minutes, seconds}, ":")
		} else {
			militaryHour := 12 + hourVal
			return strings.Join([]string{strconv.Itoa(militaryHour), minutes, seconds}, ":")
		}
	} else if hourVal == 12 {
		//12 AM
		return strings.Join([]string{"00", minutes, seconds}, ":")
	}
	return strings.Join([]string{hours, minutes, seconds}, ":")
}
