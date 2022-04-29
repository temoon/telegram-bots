package bots

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/temoon/telegram-bots/config"
)

func InitLog() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		PrettyPrint:     config.IsTestEnvironment(),
		DataKey:         "meta",
	})

	log.SetOutput(os.Stdout)

	if config.IsTestEnvironment() {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
