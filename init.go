package bots

import (
	"github.com/rs/zerolog"

	"github.com/temoon/telegram-bots/config"
)

//goland:noinspection GoUnusedExportedFunction
func InitLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if config.IsTestEnvironment() {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
