package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/btncallback"
	"github.com/rachmanzz/nextjob/setup"
)

func HandleCallbackQuery(update tgbotapi.Update) {
	var bot = setup.BOT

	bot.Request(tgbotapi.DeleteMessageConfig{
		MessageID: update.CallbackQuery.Message.MessageID,
		ChatID:    update.CallbackQuery.Message.Chat.ID,
	})

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}

	for _, v := range btncallback.RegisteredQuery {
		if callback.Text == v.Query {
			v.CallbackAction(*update.CallbackQuery.Message)
		}
	}

}
