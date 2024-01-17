package config

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	ydbCloud "github.com/yandex-cloud/ydb-go-sdk/v2"
	ydbPlatform "github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
)

const DefaultDialTimeout = time.Second * 10
const DefaultTransportTimeout = time.Millisecond * 500
const DefaultOperationTimeout = time.Millisecond * 400
const DefaultCancelAfter = time.Millisecond * 300

func GetYdbConnectionString() string {
	return "grpc://" + GetYdbEndpoint() + "/?database=" + GetYdbDatabase()
}

func GetYdbOpts() []ydbPlatform.Option {
	opts := []ydbPlatform.Option{
		ydbPlatform.WithDialTimeout(GetYdbDialTimeout()),
		ydbPlatform.WithEndpoint(GetYdbEndpoint()),
		ydbPlatform.WithDatabase(GetYdbDatabase()),
		yc.WithInternalCA(),
	}

	if true {
		opts = append(opts, ydbPlatform.WithAnonymousCredentials())
	} else {
		if IsYandexServerlessFunction() {
			opts = append(opts, yc.WithMetadataCredentials())
		} else {
			opts = append(opts, yc.WithServiceAccountKeyFileCredentials(GetYdbServiceAccountKeyFileCredentials()))
		}
	}

	return opts
}

//goland:noinspection GoUnusedExportedFunction
func GetYdbTimeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	transportTimeoutCtx, cancel := context.WithTimeout(ctx, GetYdbTransportTimeout())
	operationTimeoutCtx := ydbCloud.WithOperationTimeout(transportTimeoutCtx, GetYdbOperationTimeout())
	cancelAfterCtx := ydbCloud.WithOperationCancelAfter(operationTimeoutCtx, GetYdbCancelAfter())

	return cancelAfterCtx, cancel
}

func GetYdbDialTimeout() time.Duration {
	if value := os.Getenv("YDB_DIAL_TIMEOUT"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.WithError(err).Fatal("invalid Yandex DB dial timeout")
		} else {
			return timeout
		}
	}

	return DefaultDialTimeout
}

func GetYdbTransportTimeout() time.Duration {
	if value := os.Getenv("YDB_TRANSPORT_TIMEOUT"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.WithError(err).Fatal("invalid Yandex DB transport timeout")
		} else {
			return timeout
		}
	}

	return DefaultTransportTimeout
}

func GetYdbOperationTimeout() time.Duration {
	if value := os.Getenv("YDB_OPERATION_TIMEOUT"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.WithError(err).Fatal("invalid Yandex DB operation timeout")
		} else {
			return timeout
		}
	}

	return DefaultOperationTimeout
}

func GetYdbCancelAfter() time.Duration {
	if value := os.Getenv("YDB_CANCEL_AFTER"); value != "" {
		if timeout, err := time.ParseDuration(value); err != nil {
			log.WithError(err).Fatal("invalid Yandex DB cancel after")
		} else {
			return timeout
		}
	}

	return DefaultCancelAfter
}

func GetYdbEndpoint() (value string) {
	if value = os.Getenv("YDB_ENDPOINT"); value == "" {
		log.Fatal("Yandex DB endpoint required")
	}

	return
}

func GetYdbDatabase() (value string) {
	if value = os.Getenv("YDB_DATABASE"); value == "" {
		log.Fatal("Yandex DB database required")
	}

	return
}

func GetYdbServiceAccountKeyFileCredentials() (value string) {
	if value = os.Getenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS"); value == "" {
		log.Fatal("Yandex DB service account key file credentials required")
	}

	return
}
