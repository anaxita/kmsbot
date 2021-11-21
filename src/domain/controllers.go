package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net"
	"strings"
)

const (
	commandStart = "start"
	commandHelp  = "help"
)

const (
	btnMikrotik      = "mikrotik"
	btnAddIP         = "add ip"
	btnRemoveIP      = "remove ip"
	btnToStart       = "to start"
	btnToChatsList   = "chats list"
	btnToAdminsList  = "admins list"
	btnDeclinedAddIp = "decline add ip"
)

func (c *Core) callbackController(callbackQuery *tgbotapi.CallbackQuery) {
	var data = callbackQuery.Data

	switch data {
	default:
		log.Println("unknown callback data ", data)
	case btnToStart:
		c.callbackToStartHandler(callbackQuery)
	case btnAddIP:
		c.callbackAddIPHandler(callbackQuery)
	case btnMikrotik:
		c.callbackToMikrotikHandler(callbackQuery)
	case btnToChatsList:
		c.callbackToChatsListHandler(callbackQuery)
	case btnToAdminsList:
		c.callbackToAdminsListHandler(callbackQuery)
	case btnDeclinedAddIp:
		c.callbackDeclineAddIPHandler(callbackQuery)
	}

}

func (c *Core) commandController(update tgbotapi.Update) {
	var command = update.Message.Command()

	switch command {
	default:
		log.Println("unknown command ", command)
	case commandStart:
		c.commandStartHandler(update)
	}

}

func (c *Core) messageController(update tgbotapi.Update) {
	var msg tgbotapi.MessageConfig
	var text = update.Message.Text

	ip, isIp := isContainIP(text)
	//ipNet, isIpNet := isContainIpNet(text)

	switch {
	default:
		log.Println("unknown message ", text)
	case isIp:
		msg = c.AskToAddIPMessageHandler(update, ip)
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}

func isContainIP(text string) (string, bool) {
	split := strings.Split(text, " ")

	for _, s := range split {
		ip := net.ParseIP(s)
		if ip != nil {
			return ip.String(), true
		}
	}

	return "", false
}

func isContainIpNet(text string) (string, bool) {
	split := strings.Split(text, " ")

	for _, s := range split {
		_, ipNetwork, err := net.ParseCIDR(s)
		if err != nil {
			return ipNetwork.String(), true
		}
	}

	return "", false
}
