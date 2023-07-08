package bots

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/temoon/telegram-bots/config"
)

//goland:noinspection GoUnusedGlobalVariable
var Location *time.Location

func init() {
	var err error
	if Location, err = time.LoadLocation(config.GetTimezone()); err != nil {
		log.Fatal().Err(err).Msg("load location")
	}
}
