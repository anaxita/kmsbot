package domain

import (
	"kmsbot/service"
	"log"
)

type Core struct {
	bot      *service.Bot
	store    *service.Store
	mikrotik *service.Mikrotik
}

func NewCore(bot *service.Bot, store *service.Store, mikrotik *service.Mikrotik) *Core {
	return &Core{bot: bot, store: store, mikrotik: mikrotik}
}

func (c *Core) Start() {
	updates := c.bot.GetUpdatesChan(c.bot.Config)

	for update := range updates {
		if update.CallbackQuery != nil {
			log.Println("[INFO] Got a callback")
			c.callbackController(update.CallbackQuery)

			continue
		}

		if update.Message == nil {
			log.Println("[INFO] Got a nil message")

			continue
		}

		if update.Message.Chat.Type == "private" {
			return
		}

		if update.Message.IsCommand() {
			log.Println("[INFO] Got a command")

			c.commandController(update)

			continue
		}

		log.Println("[INFO] Got a message")

		c.messageController(update)
	}
}
