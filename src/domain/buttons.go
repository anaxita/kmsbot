package domain

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var startAdminKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Mikrotik", btnMikrotik),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Чаты", btnToChatsList),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Администраторы", btnToAdminsList),
	),
)

var wlAdminKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить IP", btnAddIP),
		tgbotapi.NewInlineKeyboardButtonData("Удалить IP", btnAddIP),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", btnToStart),
	),
)

var chatsListAdminKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить чат", btnAddIP),
		tgbotapi.NewInlineKeyboardButtonData("Удалить чат", btnAddIP),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Список чатов", btnAddIP),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", btnToStart),
	),
)

var adminsListAdminKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить админа", btnAddIP),
		tgbotapi.NewInlineKeyboardButtonData("Удалить удамина", btnAddIP),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Список админов", btnAddIP),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", btnToStart),
	),
)

var askToAddIPClientKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да", btnAddIP),
		tgbotapi.NewInlineKeyboardButtonData("Нет", btnDeclinedAddIp),
	),
)
