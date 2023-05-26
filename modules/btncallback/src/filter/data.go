package filter

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

func FilterData(msg tgbotapi.Message) {

	var bot = setup.BOT
	var redis = setup.REDIS
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")
	keyID := fmt.Sprintf("filter_%d", msg.Chat.ID)
	msgData := ""

	var val string
	var err error

	if val, err = redis.Get(ctx, keyID).Result(); err == nil {
		inputfilter := types.FilterRefObject{}
		inputfilter.ToObject(val)

		if inputfilter.Locations != nil && len(inputfilter.Locations) >= 1 {
			location := ""
			for _, item := range inputfilter.Locations {
				if location != "" {
					location += "; " + item
				} else {
					location = item
				}
			}
			if location != "" {
				msgData += "location: " + location + "\n"
			}
		}
		if inputfilter.Currency != "" {
			msgData += "currency: " + inputfilter.Currency + "\n"
		}
		if inputfilter.Type != "" {
			msgData += "type: " + inputfilter.Type + "\n"
		}
		if inputfilter.SalaryFrom != "" {
			salary := inputfilter.SalaryFrom
			if inputfilter.SalaryTo != "" {
				salary += " - " + inputfilter.SalaryTo
			}
			msgData += "salary: " + salary + "\n"
		}

	}
	if msgData != "" {
		msgConfig.Text = msgData
	} else {
		msgConfig.Text = "no data..."
	}

	if _, err := bot.Send(msgConfig); err != nil {
		log.Println(err.Error())
	}
}
