package bots

import (
	"encoding/base64"
	"errors"

	"github.com/vmihailenco/msgpack/v5"
)

const (
	FlagNewMessage Flags = 1 << iota
	FlagConfirmed
)

type CommandCode uint8
type Flags uint8

type CallbackData struct {
	_msgpack struct{} `msgpack:",as_array"`

	Command  CommandCode
	Flags    Flags
	Payload  []byte
	ReturnTo *CallbackData
}

func DecodeCallbackData(payload string) (callbackData *CallbackData, err error) {
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

func (d *CallbackData) DecodePayload(payload interface{}) (err error) {
	return msgpack.Unmarshal(d.Payload, payload)
}

func (d *CallbackData) EncodePayload(payload interface{}) (err error) {
	d.Payload, err = msgpack.Marshal(payload)
	return
}

func (d *CallbackData) String() string {
	data, err := msgpack.Marshal(d)
	if err != nil {
		return ""
	}

	return base64.RawStdEncoding.EncodeToString(data)
}

func (d *CallbackData) NewMessage() bool {
	return d.Flags&FlagNewMessage != 0
}

func (d *CallbackData) Confirmed() bool {
	return d.Flags&FlagConfirmed != 0
}
