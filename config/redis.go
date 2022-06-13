package config

import (
	"os"
	"strings"
)

func GetRedisAddress() []string {
	var value string
	if value = os.Getenv("REDIS_ADDRESS"); value == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}
