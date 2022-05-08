module github.com/temoon/telegram-bots

go 1.16

replace github.com/temoon/telegram-bots-api => ../telegram-bots-api

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/temoon/telegram-bots-api v1.600.4
	github.com/vmihailenco/msgpack/v5 v5.3.5
	github.com/yandex-cloud/ydb-go-sdk/v2 v2.12.0 // indirect
	github.com/ydb-platform/ydb-go-sdk/v3 v3.24.2 // indirect
	github.com/ydb-platform/ydb-go-yc v0.8.2 // indirect
	golang.org/x/sys v0.0.0-20220429121018-84afa8d3f7b3 // indirect
)
