package src

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

var filterKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("location", "--filter-location"),
		tgbotapi.NewInlineKeyboardButtonData("type", "--filter-type"),
		tgbotapi.NewInlineKeyboardButtonData("currency", "--filter-currency"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("salary", "--filter-salary"),
		tgbotapi.NewInlineKeyboardButtonData("inline arg", "--filter-inline-arg"),
	),
)

func FilterCommand(msg *tgbotapi.MessageConfig, opts ...any) {
	var bot = setup.BOT

	msg.Text = "please select the filter button below"
	msg.ReplyMarkup = filterKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.Println(err.Error())
	}

}
