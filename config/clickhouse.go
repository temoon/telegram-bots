package config

import (
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	log "github.com/sirupsen/logrus"
)

func GetClickHouseOpts() *clickhouse.Options {
	return &clickhouse.Options{
		Addr: []string{os.Getenv("CLICKHOUSE_ADDRESS")},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USER"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
		Debug: true,
		Debugf: func(format string, v ...interface{}) {
			log.Debugf(format, v)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:      time.Duration(10) * time.Second,
		MaxOpenConns:     5,
		MaxIdleConns:     5,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		BlockBufferSize:  10,
	}
}
