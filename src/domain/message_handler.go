package domain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Core) AskToAddIPMessageHandler(update tgbotapi.Update, ip string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Здравствуйте, я - робот КМ Системс.\nПодскажите, нужно ли добавить  адрес %s в белый список?", ip))
	msg.ReplyMarkup = askToAddIPClientKeyboard

	return msg
}
