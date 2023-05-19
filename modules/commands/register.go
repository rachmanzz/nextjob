package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands/src"
)

type RegisteredCommandType struct {
	CommandName   string
	CommandAction func(msg *tgbotapi.MessageConfig, opt ...any)
}

var RegisteredCommand []RegisteredCommandType

func init() {
	RegisteredCommand = []RegisteredCommandType{
		{
			CommandName:   "start",
			CommandAction: src.StartCommand,
		},
		{
			CommandName:   "how",
			CommandAction: src.HowCommand,
		},
	}
}
