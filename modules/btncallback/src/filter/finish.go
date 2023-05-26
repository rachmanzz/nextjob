package filter

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

func FilterBtnFinish(msg tgbotapi.Message) {

	var bot = setup.BOT
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")

	msgConfig.Text = "filter saved for 8 hours..."

	if _, err := bot.Send(msgConfig); err != nil {
		log.Println(err.Error())
	}
}
