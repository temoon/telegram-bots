package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/temoon/telegram-bots-api"
)

func IsYandexServerlessFunction() bool {
	return os.Getenv("_HANDLER") != ""
}

func IsHttpServer() bool {
	return !IsTestEnvironment() && os.Getenv("HTTP_ADDRESS") != ""
}

func IsTestEnvironment() bool {
	return GetBotEnvironment() == telegram.EnvTest
}

func EnvToInt64Slice(key string) []int64 {
	var value string
	if value = os.Getenv(key); value == "" {
		return []int64{}
	}

	strItems := strings.Split(value, ",")
	items := make([]int64, 0, len(strItems))

	var item int64
	var err error
	for _, strItem := range strItems {
		strItem = strings.TrimSpace(strItem)
		if item, err = strconv.ParseInt(strItem, 10, 64); err != nil {
			continue
		}
		items = append(items, item)
	}

	return items
}
