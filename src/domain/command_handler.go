package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Core) commandStartHandler(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите команду.")
	msg.ReplyMarkup = startAdminKeyboard
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func (c *Core) callbackAddIPHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.MessageConfig

	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.Text = "Callback data: " + callbackQuery.Data

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		msg.Text = "Техническая ошибка, попробуйте снова"
	}

	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func (c *Core) callbackToStartHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.EditMessageReplyMarkupConfig

	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.MessageID = callbackQuery.Message.MessageID
	msg.ReplyMarkup = &startAdminKeyboard

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		log.Println("callback request ", err)

	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}
