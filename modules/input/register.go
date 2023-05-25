package input

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/input/src/filter"
)

type RegisterInputCommandType struct {
	CommandName   string
	CommandAction func(refData string, botMsg *tgbotapi.Message)
}

var RegisteredInputCommand []RegisterInputCommandType

func init() {
	RegisteredInputCommand = []RegisterInputCommandType{
		{
			CommandName:   "/filter",
			CommandAction: filter.FilterReferance,
		},
	}
}
