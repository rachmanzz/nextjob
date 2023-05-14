package src

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands/options"
	"github.com/rachmanzz/nextjob/setup"
)

func StartCommand(msg *tgbotapi.MessageConfig, opts ...any) {
	var bot = setup.BOT
	opt := options.GetOpt(opts...)

	str := fmt.Sprintf("Hi %s\nWelcome to NextJob Seeker. Find the best job and apply.", opt.Name)
	msg.Text = str
	if _, err := bot.Send(msg); err != nil {
		log.Println(err.Error())
	}
}
