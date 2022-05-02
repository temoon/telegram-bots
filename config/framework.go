package config

import (
	"os"

	"github.com/temoon/telegram-bots-api"
)

func IsServerlessFunction() bool {
	return os.Getenv("_HANDLER") != ""
}

func IsHttpServer() bool {
	return !IsTestEnvironment() && os.Getenv("HTTP_ADDRESS") != ""
}

func IsTestEnvironment() bool {
	return GetBotEnvironment() == telegram.EnvTest
}
