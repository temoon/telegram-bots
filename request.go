package bots

import (
	"strings"

	"github.com/temoon/telegram-bots-api"

	. "github.com/temoon/telegram-bots/helpers"
)

type Request struct {
	BotUsername     string
	Command         string
	Payload         string
	ChatId          telegram.ChatId
	Chat            *telegram.Chat
	UserId          int64
	User            *telegram.User
	MessageId       int64
	Message         *telegram.Message
	CallbackQueryId string
	CallbackData    *CallbackData
}

func (req *Request) HasContact() bool {
	return req.Message != nil && req.Message.Contact != nil
}

func (req *Request) IsCommand(str string, code CommandCode) bool {
	return req.Command == str ||
		(req.ChatId.GetId() < 0 && req.BotUsername != "" && req.Command == str+"@"+req.BotUsername) ||
		(req.CallbackData != nil && req.CallbackData.Command == code)
}

func (req *Request) GetCallbackDataPayload() interface{} {
	if req.CallbackData != nil {
		return req.CallbackData.Payload
	}

	return nil
}

func (req *Request) NewMessage() bool {
	return req.CallbackData == nil || req.CallbackData.NewMessage() || req.MessageId == 0
}

func (req *Request) Confirmed() bool {
	return req.CallbackData == nil || req.CallbackData.Confirmed()
}

func ParseMessage(message *telegram.Message) (req *Request, err error) {
	command, payload := parseText(StringOrEmpty(message.Text))

	req = &Request{
		Command:   command,
		Payload:   payload,
		ChatId:    ChatId(message.Chat.Id),
		Chat:      &message.Chat,
		MessageId: message.MessageId,
		Message:   message,
	}

	if message.From != nil {
		req.UserId = message.From.Id
		req.User = message.From
	}

	return
}

func ParseCallbackQuery(callbackQuery *telegram.CallbackQuery) (req *Request, err error) {
	payload := StringOrEmpty(callbackQuery.Data)

	var callbackData *CallbackData
	if callbackData, err = DecodeCallbackData(payload); err != nil {
		return
	}

	req = &Request{
		Payload:         payload,
		UserId:          callbackQuery.From.Id,
		User:            &callbackQuery.From,
		CallbackQueryId: callbackQuery.Id,
		CallbackData:    callbackData,
	}

	if callbackQuery.Message != nil {
		switch message := callbackQuery.Message.(type) {
		case telegram.Message:
			req.ChatId = ChatId(message.Chat.Id)
			req.Chat = &message.Chat
			req.MessageId = message.MessageId
			req.Message = &message
		case telegram.InaccessibleMessage:
			req.ChatId = ChatId(message.Chat.Id)
			req.Chat = &message.Chat
		}
	}

	return
}

func parseText(text string) (command, payload string) {
	if text == "" || text[0] != '/' {
		return "", text
	}

	parts := strings.SplitN(text, " ", 2)
	if len(parts) > 1 {
		command = parts[0]
		payload = parts[1]
	} else {
		command = parts[0]
	}

	return
}
