package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Core) commandStartAdminHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите команду.")
	msg.ReplyMarkup = startAdminKeyboard

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}
