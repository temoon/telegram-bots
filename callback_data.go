package bots

import (
	"encoding/base64"
	"errors"

	"github.com/vmihailenco/msgpack/v5"
)

type CallbackData struct {
	_msgpack struct{} `msgpack:",as_array"`

	Command            int8
	Payload            interface{}
	NewMessageRequired bool
}

func ParseCallbackData(payload string) (callbackData *CallbackData, err error) {
	if len(payload) == 0 {
		return nil, errors.New("no callback data")
	}

	var data []byte
	if data, err = base64.RawStdEncoding.DecodeString(payload); err != nil {
		return
	}

	if err = msgpack.Unmarshal(data, &callbackData); err != nil {
		return
	}

	return
}

func (d *CallbackData) String() string {
	data, err := msgpack.Marshal(d)
	if err != nil {
		return ""
	}

	return base64.RawStdEncoding.EncodeToString(data)
}
