package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/telegram-bots-api"
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
	for _, disallowedUserId := range GetBotDisallowedUsers() {
		if disallowedUserId == userId {
			return false
		}
	}

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
	return EnvToInt64Slice("BOT_ALLOWED_USERS")
}

func GetBotDisallowedUsers() []int64 {
	return EnvToInt64Slice("BOT_DISALLOWED_USERS")
}
