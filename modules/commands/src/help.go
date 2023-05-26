package src

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands/options"
	"github.com/rachmanzz/nextjob/setup"
)

var helpKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("docs", "https://rachmanzz.github.io/nextjob/"),
	),
)

func HelpCommand(msg *tgbotapi.MessageConfig, opts ...any) {
	var bot = setup.BOT
	opt := options.GetOpt(opts...)

	str := fmt.Sprintf("Hi %s!\nPlease send command /how to find how to start the bot, but if not informate well. \nplease tap/click button bellow.", opt.Name)
	msg.Text = str
	msg.ReplyMarkup = helpKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.Println(err.Error())
	}
}
