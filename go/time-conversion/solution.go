package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	isAm := strings.HasSuffix(s, "AM")
	timeParts := strings.Split(s, ":")
	hour := timeParts[0]

	if isAm {
		if hour == "12" {
			timeParts[0] = "00"
		}
	} else {
		intHour, _ := strconv.Atoi(hour)
		if intHour < 12 {
			intHour = intHour + 12
		}
		timeParts[0] = strconv.Itoa(intHour)
	}

	return fmt.Sprintf("%s:%s:%s", timeParts[0], timeParts[1], timeParts[2][:2])
}
