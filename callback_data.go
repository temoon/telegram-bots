package bots

import (
	"encoding/base64"
	"errors"

	"github.com/vmihailenco/msgpack/v5"
)

type CallbackData struct {
	_msgpack struct{} `msgpack:",as_array"`

	Command int8
	ReplyTo bool // TODO маркер, нужно ли отвечать на сообщение или отрисовать новое
	Payload interface{}
	History *CallbackData
}

func ParseCallbackData(data []byte) (callbackData *CallbackData, err error) {
	if len(data) == 0 {
		return nil, errors.New("no callback data")
	}

	if err = msgpack.Unmarshal(data, &callbackData); err != nil {
		return
	}

	return
}

func (p *CallbackData) String() string {
	data, err := msgpack.Marshal(p)
	if err != nil {
		return ""
	}

	return base64.RawStdEncoding.EncodeToString(data)
}
