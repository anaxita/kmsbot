package domain

import (
	"kmsbot/service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
	defer func() {
		err := recover()
		if err != nil {
			log.Println("recovered: ", err)
		}
	}()

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
	defer func() {
		err := recover()
		if err != nil {
			log.Println("recovered: ", err)
		}
	}()

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
	defer func() {
		err := recover()
		if err != nil {
			log.Println("recovered: ", err)
		}
	}()

	var text = update.Message.Text
	ip, isIp := isContainIP(text)

	isAdminChat := c.isAdminChat(update.Message.Chat.ID)

	if isAdminChat {
		ipNet, isNet := isContainIpNet(text)

		switch {
		case isNet:
			c.AskToAddIPMessageHandler(update, ipNet)
		case isIp:
			c.AskToAddIPMessageHandler(update, ip)
		}

		return
	}

	switch {
	case len(update.Message.NewChatMembers) > 0:
		c.SendGreetingMessageHandler(update)
	case isIp:
		c.AskToAddIPMessageHandler(update, ip)
	}
}

func (c *Core) eventController(update tgbotapi.Update) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("recovered: ", err)
		}
	}()

	chat, err := c.store.ChatByChatID(update.MyChatMember.Chat.ID)
	if err != nil {
		log.Println("???????????? ?????????????? ???????? ???? ????", err)
		return
	}

	if chat == nil &&
		update.MyChatMember.NewChatMember.Status == "member" &&
		(update.MyChatMember.From.UserName == "anaxita" ||
			update.MyChatMember.From.UserName == "Mishagl" ||
			update.MyChatMember.From.UserName == "KM_SYSTEM") {

		log.Println("?????? ???? ???????????? ?? ????")

		chat = &service.Chat{
			Title:  update.MyChatMember.Chat.Title,
			ChatID: update.MyChatMember.Chat.ID,
			Role:   service.RoleClient,
		}

		log.Println("???????????????? ?????????? ?????? ?? ????")

		err := c.store.CreateChat(*chat)
		if err != nil {
			log.Println("???????????? ???????????????????? ???????? ?? ????", err)

			return
		}
	}

	if chat == nil {
		log.Println("???????? ???????? ?? ????, ?????????????????? ?????????????????? ??????????????????")

		msg := tgbotapi.LeaveChatConfig{
			ChatID: update.MyChatMember.Chat.ID,
		}

		log.Println("?????????????? ???????????? ??????")

		_, err := c.bot.Request(msg)
		if err != nil {
			log.Println("???????????? ???????????? ???? ????????", err)
		}

		return
	}
}
