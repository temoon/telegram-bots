package config

import (
	"os"
)

const DefaultTimezone = "Europe/Moscow"

func GetTimezone() (value string) {
	if value = os.Getenv("TZ"); value != "" {
		return
	}

	return DefaultTimezone
}
