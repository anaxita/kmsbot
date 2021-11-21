package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Core) callbackToMikrotikHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.EditMessageReplyMarkupConfig

	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.MessageID = callbackQuery.Message.MessageID
	msg.ReplyMarkup = &wlAdminKeyboard

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		log.Println("callback request ", err)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func (c *Core) callbackToChatsListHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.EditMessageReplyMarkupConfig

	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.MessageID = callbackQuery.Message.MessageID
	msg.ReplyMarkup = &chatsListAdminKeyboard

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		log.Println("callback request ", err)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func (c *Core) callbackToAdminsListHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.EditMessageReplyMarkupConfig

	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.MessageID = callbackQuery.Message.MessageID
	msg.ReplyMarkup = &adminsListAdminKeyboard

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		log.Println("callback request ", err)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func (c *Core) callbackDeclineAddIPHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var msg tgbotapi.DeleteMessageConfig
	msg.ChatID = callbackQuery.Message.Chat.ID
	msg.MessageID = callbackQuery.Message.MessageID
	callbackQuery.Message.ReplyMarkup = nil
	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		log.Println("callback request ", err)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}
