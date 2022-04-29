package helpers

import (
	"context"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/temoon/go-telegram-bots-api"
	. "github.com/temoon/go-telegram-bots-api/helpers"
	"github.com/temoon/go-telegram-bots-api/requests"
)

func GetUserName(user *telegram.User) (name string) {
	if name = strings.TrimSpace(user.FirstName + " " + StringOrEmpty(user.LastName)); name != "" {
		return
	}

	if name = StringOrEmpty(user.Username); name != "" {
		return
	}

	return strconv.FormatInt(user.Id, 10)
}

func SendInfoMessage(ctx context.Context, bot *telegram.Bot, chatId int64, text string) (err error) {
	text = "💬 " + text

	message := &requests.SendMessage{
		ChatId: chatId,
		Text:   text,
	}

	if _, err = message.Call(ctx, bot); err != nil {
		return
	}

	return
}

func SendErrorMessage(ctx context.Context, bot *telegram.Bot, chatId int64) {
	text := "*⚠ Ошибка!* Что-то пошло не так, но мы уже чиним."

	message := &requests.SendMessage{
		ChatId:    chatId,
		ParseMode: String(telegram.ParseModeMarkdown),
		Text:      text,
	}

	if _, err := message.Call(ctx, bot); err != nil {
		log.WithError(err).Error("can't send error message")
		return
	}

	return
}
