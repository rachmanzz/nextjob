package filter

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands/src"
	"github.com/rachmanzz/nextjob/setup"
)

func FilterOpen(msg tgbotapi.Message) {

	var bot = setup.BOT
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")

	msgConfig.Text = "please select the filter button below"
	msgConfig.ReplyMarkup = src.FilterKeyboard

	if _, err := bot.Send(msgConfig); err != nil {
		log.Println(err.Error())
	}
}
