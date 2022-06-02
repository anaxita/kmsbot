package domain

import (
	"kmsbot/service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Core struct {
	bot       *service.Bot
	store     *service.Store
	mikrotik  *service.Mikrotik
	mikrotik2 *service.Mikrotik
}

func NewCore(bot *service.Bot, store *service.Store, mikrotik *service.Mikrotik, mikrotik2 *service.Mikrotik) *Core {
	return &Core{bot: bot, store: store, mikrotik: mikrotik, mikrotik2: mikrotik2}
}

func (c *Core) Start() {
	updates := c.bot.GetUpdatesChan(c.bot.Config)

	for update := range updates {

		if update.CallbackQuery != nil {
			log.Println("[INFO] Got a callback")
			c.callbackController(update.CallbackQuery)

			continue
		}

		if update.MyChatMember != nil {
			c.eventController(update)

			continue
		}
		if update.Message == nil {
			log.Println("[INFO] Got a nil message")

			continue
		}

		if update.Message.Chat.IsPrivate() {
			_, err := c.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,
				"Извините, я не отвечаю в личные сообщение. Напишите в общий чат."))
			if err != nil {
				log.Println("send private message: ", err)
			}
			continue
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
