package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kmsbot/service"
	"log"
)

const (
	commandStart = "start"
	commandHelp  = "help"
)

const (
	btnMikrotik      = "mikrotik"
	btnAddIP         = "add ip"
	btnAcceptAddIP   = "accept adding ip"
	btnRemoveIP      = "remove ip"
	btnToStart       = "to start"
	btnToChatsList   = "chats list"
	btnToAdminsList  = "admins list"
	btnDeclinedAddIP = "decline add ip"
)

const (
	anaxitaUsername  = "anaxita"
	mishaglUsername  = "Mishagl"
	kmsControlChatID = -1001700493413
	kmsMailChatID    = -1001287143568
)

func (c *Core) callbackController(callbackQuery *tgbotapi.CallbackQuery) {
	var data = callbackQuery.Data
	isAdminChat := c.isAdminChat(callbackQuery.Message.Chat.ID)

	if isAdminChat {
		switch data {
		default:
			log.Println("unknown admin callback data", data)
		case btnToStart:
			c.callbackToStartHandler(callbackQuery)
		case btnMikrotik:
			c.callbackToMikrotikHandler(callbackQuery)
		case btnToChatsList:
			c.callbackToChatsListHandler(callbackQuery)
		case btnToAdminsList:
			c.callbackToAdminsListHandler(callbackQuery)

		case btnAcceptAddIP:
			c.callbackAddIPHandler(callbackQuery)
		case btnDeclinedAddIP:
			c.callbackDeclineAddIPHandler(callbackQuery)
		}

		return
	}

	switch data {
	default:
		log.Println("unknown client callback data", data)
	case btnDeclinedAddIP:
		c.callbackDeclineAddIPHandler(callbackQuery)
	case btnAcceptAddIP:
		c.callbackAddIPHandler(callbackQuery)
	}

}

func (c *Core) commandController(update tgbotapi.Update) {
	var command = update.Message.Command()

	isAdminChat := c.isAdminChat(update.Message.Chat.ID)

	if isAdminChat {
		switch command {
		default:
			log.Println("unknown admin command ", command)
		case commandStart:
			c.commandStartAdminHandler(update)
		}

		return
	}

	switch command {
	default:
		log.Println("unknown client command ", command)
	case commandStart:
		c.SendGreetingMessageHandler(update)
	}
}

func (c *Core) messageController(update tgbotapi.Update) {
	var text = update.Message.Text
	ip, isIp := isContainIP(text)
	ipNet, isNet := isContainIpNet(text)

	isAdminChat := c.isAdminChat(update.Message.Chat.ID)

	if isAdminChat {
		switch {
		default:
			log.Printf("[MESSAGE] %#v\n", update.Message.NewChatMembers)
		case isIp:
			c.AskToAddIPMessageHandler(update, ip)
		case isNet:
			c.AskToAddIPMessageHandler(update, ipNet)
		}

		return
	}

	switch {
	default:
		log.Printf("[MESSAGE] %s", update.Message.Text)
	case len(update.Message.NewChatMembers) > 0:
		c.SendGreetingMessageHandler(update)
	case isIp:
		c.AskToAddIPMessageHandler(update, ip)
	}
}

func (c *Core) eventController(update tgbotapi.Update) {
	chat, err := c.store.ChatByChatID(update.MyChatMember.Chat.ID)
	if err != nil {
		log.Println("Ошибка запроса чата из БД", err)
		return
	}

	if chat == nil &&
		update.ChatMember.NewChatMember.Status == "member" &&
		(update.ChatMember.From.UserName == "anaxita" ||
			update.ChatMember.From.UserName == "Mishagl") {

		log.Println("Чат не найден в БД")

		chat = &service.Chat{
			Title:  update.ChatMember.Chat.Title,
			ChatID: update.ChatMember.Chat.ID,
			Role:   service.RoleClient,
		}

		log.Println("Добавляю новый чат в БД")

		err := c.store.CreateChat(*chat)
		if err != nil {
			log.Println("Ошибка добавления чата в БД", err)

			return
		}
	}

	if chat == nil {
		log.Println("Чата нету в БД, обработка сообщений отключена")

		msg := tgbotapi.LeaveChatConfig{
			ChatID: update.ChatMember.Chat.ID,
		}

		log.Println("Покидаю данный чат")

		_, err := c.bot.Request(msg)
		if err != nil {
			log.Println("Ошибка выхода из чата", err)
		}

		return
	}
}
