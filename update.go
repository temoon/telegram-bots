package bots

import (
	"strings"

	"github.com/temoon/telegram-bots-api"
	. "github.com/temoon/telegram-bots-api/helpers"
)

type Update struct {
	Command         string
	Payload         string
	ChatId          int64
	Chat            *telegram.Chat
	UserId          int64
	User            *telegram.User
	MessageId       int32
	Message         *telegram.Message
	CallbackQueryId string
	CallbackData    *CallbackData
}

func (u *Update) HasContact() bool {
	return u.Message != nil && u.Message.Contact != nil
}

func (u *Update) CheckCommand(str string, code int8) bool {
	return u.Command == str || u.CallbackData != nil && u.CallbackData.Command == code
}

func (u *Update) NewMessageRequired() bool {
	return u.CallbackData == nil || u.CallbackData.NewMessageRequired
}

func ParseMessage(message *telegram.Message) (update *Update, err error) {
	command, payload := parseText(StringOrEmpty(message.Text))

	update = &Update{
		Command:   command,
		Payload:   payload,
		ChatId:    message.Chat.Id,
		Chat:      &message.Chat,
		MessageId: message.MessageId,
		Message:   message,
	}

	if message.From != nil {
		update.UserId = message.From.Id
		update.User = message.From
	}

	return
}

func ParseCallbackQuery(callbackQuery *telegram.CallbackQuery) (update *Update, err error) {
	payload := StringOrEmpty(callbackQuery.Data)

	var callbackData *CallbackData
	if callbackData, err = ParseCallbackData(payload); err != nil {
		return
	}

	update = &Update{
		Payload:         payload,
		UserId:          callbackQuery.From.Id,
		User:            &callbackQuery.From,
		CallbackQueryId: callbackQuery.Id,
		CallbackData:    callbackData,
	}

	if callbackQuery.Message != nil {
		update.ChatId = callbackQuery.Message.Chat.Id
		update.Chat = &callbackQuery.Message.Chat
		update.MessageId = callbackQuery.Message.MessageId
		update.Message = callbackQuery.Message
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
