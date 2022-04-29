package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/go-telegram-bots-api"
)

const DefaultBotTimeout = time.Second * 10

func GetBotToken() (value string) {
	if value = os.Getenv("BOT_TOKEN"); value == "" {
		log.Fatal("bot token required")
	}

	return
}

func GetBotTimeout() time.Duration {
	if value := os.Getenv("BOT_TIMEOUT"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.WithError(err).Fatal("invalid bot timeout")
		} else {
			return timeout
		}
	}

	return DefaultBotTimeout
}

func GetBotEnvironment() (value string) {
	value = os.Getenv("BOT_ENVIRONMENT")
	switch value {
	case telegram.EnvProduction, telegram.EnvTest:
		return
	default:
		return telegram.EnvProduction
	}
}

func IsBotUserAllowed(userId int64) bool {
	allowedUsers := GetBotAllowedUsers()
	if len(allowedUsers) == 0 {
		return true
	}

	for _, allowedUserId := range allowedUsers {
		if allowedUserId == userId {
			return true
		}
	}

	return false
}

func GetBotAllowedUsers() []int64 {
	return envToInt64Slice("BOT_ALLOWED_USERS")
}

func envToInt64Slice(key string) []int64 {
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
