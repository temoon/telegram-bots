package bots

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/temoon/telegram-bots/config"
)

//goland:noinspection GoUnusedGlobalVariable
var Location *time.Location

func init() {
	var err error
	if Location, err = time.LoadLocation(config.GetTimezone()); err != nil {
		log.WithError(err).Fatal("load location")
	}
}
