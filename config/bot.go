package config

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/temoon/telegram-bots-api"
)

const DefaultBotTimeout = time.Second * 10
const DefaultBotStaticRoot = "."

func GetBotToken() (value string) {
	if value = os.Getenv("BOT_TOKEN"); value == "" {
		log.Fatal().Msg("bot token required")
	}

	return
}

func GetBotTimeout() time.Duration {
	if value := os.Getenv("BOT_TIMEOUT"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.Fatal().Err(err).Msg("invalid bot timeout")
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

func GetBotStaticRoot() (value string) {
	if value = os.Getenv("LAMBDA_TASK_ROOT"); value != "" {
		return
	}

	if value = os.Getenv("BOT_STATIC_ROOT"); value != "" {
		return
	}

	return DefaultBotStaticRoot
}
