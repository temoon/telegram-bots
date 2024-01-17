package helpers

import (
	"github.com/temoon/telegram-bots-api"
	"strconv"
	"time"
)

const DayMonday = "Пн"
const DayTuesday = "Вт"
const DayWednesday = "Ср"
const DayThursday = "Чт"
const DayFriday = "Пт"
const DaySaturday = "Сб"
const DaySunday = "Вс"

func ChatId(chatId int64) telegram.ChatId {
	return telegram.NewChatId(chatId, "")
}

func Bool(v bool) *bool           { return &v }
func Int(v int) *int              { return &v }
func Int32(v int32) *int32        { return &v }
func Int64(v int64) *int64        { return &v }
func String(v string) *string     { return &v }
func Time(v time.Time) *time.Time { return &v }

func BoolOrFalse(ref *bool) bool {
	if ref == nil {
		return false
	}

	return *ref
}

func IntOrZero(ref *int) int {
	if ref == nil {
		return 0
	}

	return *ref
}

func Int32OrZero(ref *int32) int32 {
	if ref == nil {
		return 0
	}

	return *ref
}

func Int64OrZero(ref *int64) int64 {
	if ref == nil {
		return 0
	}

	return *ref
}

func StringOrEmpty(ref *string) string {
	if ref == nil {
		return ""
	}

	return *ref
}

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

func Url(caption, url string) string {
	return "[" + caption + "](" + url + ")"
}

func UserUrl(userId int64) string {
	return "tg://user?id=" + strconv.FormatInt(userId, 10)
}
