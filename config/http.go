package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func GetHttpAddress() (value string) {
	if value = os.Getenv("HTTP_ADDRESS"); value == "" {
		log.Fatal("http address required")
	}

	return
}

func GetHttpEndpoint() (value string) {
	if value = os.Getenv("HTTP_ENDPOINT"); value == "" {
		log.Fatal("http endpoint required")
	}

	return
}

func GetHttpCertFile() string {
	return os.Getenv("HTTP_CERT_FILE")
}

func GetHttpCertKey() string {
	return os.Getenv("HTTP_CERT_KEY")
}

func GetHttpServerName() string {
	return os.Getenv("HTTP_SERVER_NAME")
}
