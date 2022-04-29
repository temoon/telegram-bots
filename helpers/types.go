package helpers

import (
	"time"
)

const DayMonday = "Пн"
const DayTuesday = "Вт"
const DayWednesday = "Ср"
const DayThursday = "Чт"
const DayFriday = "Пт"
const DaySaturday = "Сб"
const DaySunday = "Вс"

func TimeOrNow(ref *time.Time) time.Time {
	if ref == nil {
		return time.Now()
	}

	return *ref
}

func Weekday(datetime time.Time) (day string) {
	switch datetime.Weekday() {
	case time.Monday:
		day = DayMonday
	case time.Tuesday:
		day = DayTuesday
	case time.Wednesday:
		day = DayWednesday
	case time.Thursday:
		day = DayThursday
	case time.Friday:
		day = DayFriday
	case time.Saturday:
		day = DaySaturday
	case time.Sunday:
		day = DaySunday
	}

	return
}
