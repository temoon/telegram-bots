package config

import (
	"os"
	"strings"
)

func GetKafkaSeeds() []string {
	var value string
	if value = os.Getenv("KAFKA_SEEDS"); value == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}
