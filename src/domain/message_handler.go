package domain

import (
	"fmt"
	"kmsbot/service"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Core) AskToAddIPMessageHandler(update tgbotapi.Update, ip string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
		`Здравствуйте, я - робот КМ Системс.
Нужно ли добавить %s в белый список?`,
		ip))

	msg.ReplyMarkup = askToAddIPClientKeyboard
	msg.ReplyToMessageID = update.Message.MessageID

	resp, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}

	c.store.Cache[resp.MessageID] = &service.IPMessage{IP: ip, MsgID: update.Message.MessageID}

}

func (c *Core) SendGreetingMessageHandler(update tgbotapi.Update) {
	text := fmt.Sprintf(`
Здравствуйте, я робот КМ Системс.

1. Данный чат используется для технической поддержки вашей компании.

2. Меня можно использовать для быстрого добавления IP адреса в белый список.
Просто отправьте необходимый IP в чат и я добавлю его.
`)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("[ERROR] send a message: ", err)
	}
}
