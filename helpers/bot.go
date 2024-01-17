package helpers

import (
	"context"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/temoon/telegram-bots-api"
	"github.com/temoon/telegram-bots-api/requests"

	"github.com/temoon/telegram-bots/payloads"
	. "github.com/temoon/telegram-bots/vars"
)

type ListItem interface {
	GetText() string
	GetCallbackData() *string
}

//goland:noinspection GoUnusedExportedFunction
func GetUserName(user *telegram.User) (name string) {
	if user == nil {
		return
	}

	if name = strings.TrimSpace(user.FirstName + " " + StringOrEmpty(user.LastName)); name != "" {
		return
	}

	if name = StringOrEmpty(user.Username); name != "" {
		return
	}

	return strconv.FormatInt(user.Id, 10)
}

//goland:noinspection GoUnusedExportedFunction
func SendInfoMessage(ctx context.Context, bot *telegram.Bot, chatId int64, text string) (err error) {
	text = "💬 " + text

	message := &requests.SendMessage{
		ChatId: ChatId(chatId),
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
		ChatId:    ChatId(chatId),
		ParseMode: String(telegram.ParseModeMarkdown),
		Text:      text,
	}

	if _, err := message.Call(ctx, bot); err != nil {
		log.Error().Err(err).Msg("can't send error message")
		return
	}

	return
}

//goland:noinspection GoUnusedExportedFunction
func GetListKeyboard(items []ListItem, navPayload payloads.PayloadWithOffset, offset, limit, total int) (keyboard *telegram.InlineKeyboardMarkup) {
	if len(items) == 0 || offset < 0 || limit <= 0 || total <= 0 {
		return
	}

	buttons := make([][]telegram.InlineKeyboardButton, 0, len(items)+1)
	for _, item := range items {
		buttons = append(buttons, []telegram.InlineKeyboardButton{
			{
				Text:         item.GetText(),
				CallbackData: item.GetCallbackData(),
			},
		})
	}

	nav := make([]telegram.InlineKeyboardButton, 0, 2)
	if offset > 0 {
		nav = append(nav, telegram.InlineKeyboardButton{
			Text:         TextPrev,
			CallbackData: String(navPayload.WithOffset(Max(0, offset-limit)).String()),
		})
	}
	if offset+limit < total {
		nav = append(nav, telegram.InlineKeyboardButton{
			Text:         TextNext,
			CallbackData: String(navPayload.WithOffset(Min(total, offset+limit)).String()),
		})
	}
	if len(nav) != 0 {
		buttons = append(buttons, nav)
	}

	return &telegram.InlineKeyboardMarkup{
		InlineKeyboard: buttons,
	}
}
