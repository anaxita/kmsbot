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

	ipMessage, ok := c.store.Messages[msgID]
	if !ok {
		msg.Text = "Время истекло, введите заново IP"
	}

	msg.ReplyToMessageID = ipMessage.MessageID()

	if ok {

		comment := fmt.Sprintf("BOT %s | %s %s", chatTitle, firstName, lastName)

		err := c.mikrotik.AddIP(ipMessage.IP4(), Translit(comment))
		if err != nil {
			if err.Error() == "from RouterOS device: failure: already have such entry" {
				msg.Text = "Данный IP уже находится в белом списке."
			} else {
				msg.Text = "Ошибка добавления IP: " + err.Error()
			}

		} else {
			if chatTitle == "" {
				chatTitle = "Личные сообщения"
			}

			c.SendNotification(fmt.Sprintf("Chat: %s\nUser: @%s %s %s\nAction: Добавил IP %s", chatTitle, username, firstName, lastName, ipMessage.IP4()))
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
