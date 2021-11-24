package domain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"kmsbot/service"
	"log"
	"net"
	"strings"
)

func (c *Core) sendNotification(text string) {
	msg := tgbotapi.NewMessage(kmsMailChatID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println("Can't send a message", err)
	}
}

func (c *Core) isAdminChat(chatID int64) bool {
	chat, err := c.store.ChatByChatID(chatID)
	if err != nil {
		log.Println("[ERROR] get chat by id", err)
		return false
	}

	if chat == nil {
		return false
	}

	return chat.Role == service.RoleAdmin
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

var translitMap = map[string]string{
	"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e",
	"ё": "yo", "ж": "zh", "з": "z", "и": "i", "й": "y", "к": "k",
	"л": "l", "м": "m", "н": "n", "о": "o", "п": "p", "р": "r",
	"с": "s", "т": "t", "у": "u", "ф": "f", "х": "kh", "ц": "ts",
	"ч": "ch", "ш": "sh", "щ": "sch", "ъ": "", "ы": "i",
	"ь": "", "э": "e", "ю": "yu", "я": "ya",
	"А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E",
	"Ё": "YO", "Ж": "ZH", "З": "Z", "И": "I", "Й": "Y", "К": "K",
	"Л": "L", "М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R",
	"С": "S", "Т": "T", "У": "U", "Ф": "F", "Х": "KH", "Ц": "TS",
	"Ч": "CH", "Ш": "SH", "Щ": "SCH", "Ъ": "", "Ы": "I",
	"Ь": "", "Э": "E", "Ю": "YU", "Я": "YA",
}

func Translit(text string) string {
	var result string

	for _, r := range text {
		char, ok := translitMap[string(r)]
		if !ok {
			result += string(r)

			continue
		}

		result += char
	}

	return result
}
