package domain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

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

func (c *Core) callbackAddIPHandler(callbackQuery *tgbotapi.CallbackQuery) {
	var (
		msg tgbotapi.MessageConfig

		chatTitle = callbackQuery.Message.Chat.Title
		username  = callbackQuery.From.UserName
		firstName = callbackQuery.From.FirstName
		lastName  = callbackQuery.From.LastName
		chatID    = callbackQuery.Message.Chat.ID
		msgID     = callbackQuery.Message.MessageID
	)

	msg.ChatID = chatID
	msg.Text = "IP успешно добавлен!"

	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := c.bot.Request(callback); err != nil {
		msg.Text = "Техническая ошибка, попробуйте снова"
		c.bot.Send(msg)

		return
	}

	ip, ok := c.store.Messages[msgID]
	if !ok {
		msg.Text = "Время истекло, введите заново IP"
	}

	if ok {
		comment := fmt.Sprintf("BOT %s | %s %s", chatTitle, firstName, lastName)

		err := c.mikrotik.AddIP(ip.Data(), Translit(comment))
		if err != nil {
			msg.Text = "Ошибка добавления IP: " + err.Error()
		} else {
			if chatTitle == "" {
				chatTitle = "Личные сообщения"
			}

			c.sendNotification(fmt.Sprintf("Chat: %s\nUser: @%s %s %s\nAction: Добавил IP %s", chatTitle, username, firstName, lastName, ip.Data()))
		}
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}

	delete(c.store.Messages, msgID)

	deleteMessage := tgbotapi.NewDeleteMessage(chatID, msgID)

	_, err = c.bot.Send(deleteMessage)
	if err != nil {
		log.Println("[ERROR] delete a message: ", err)
	}
}

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

	delete(c.store.Messages, msg.MessageID)
}
