package config

import (
	"os"
)

func GetPostgresUrl() string {
	return os.Getenv("POSTGRES_URL")
}
