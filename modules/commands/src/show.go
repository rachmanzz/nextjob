package src

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

func HowCommand(msg *tgbotapi.MessageConfig, opts ...any) {
	var bot = setup.BOT
	message := fmt.Sprintf("type @%s in any message input box and wait untill placeholder is popped up, then type any job position you look for you can also hit some additional filter to specify job type, location, rate of salary, etc.. .\n\n", bot.Self.UserName)

	message += "additional filter\n"

	message += "* loc : Location, state, country\n"
	message += "* type : type of job as full-time, part-time, contract, etc... \n"
	message += "* currency : usd, sgd, idr, etc... \n"
	message += "* salary : rate of salary from 0-end rate \n\n"

	message += "for example \n\n"
	message += fmt.Sprintf("@%s backend --type=full-time --currency=usd --salary=100-10000\n\n", bot.Self.UserName)
	message += "or \n\n"

	message += fmt.Sprintf("@%s backend; frontend; android --type=full-time --currency=usd --salary=100-10000\n\n", bot.Self.UserName)

	msg.Text = message
	if _, err := bot.Send(msg); err != nil {
		log.Println(err.Error())
	}

}
