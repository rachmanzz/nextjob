package filter

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/textparse"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

var ctx = context.Background()

var filterKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("next", "--filter-open"),
		tgbotapi.NewInlineKeyboardButtonData("finish", "--filter-close"),
	),
)

func FilterReferance(ref string, botMsg *tgbotapi.Message) {
	var redis = setup.REDIS
	keyID := fmt.Sprintf("filter_%d", botMsg.Chat.ID)
	inputfilter := types.FilterRefObject{}
	hasChange := false
	if val, err := redis.Get(ctx, keyID).Result(); err == nil {
		inputfilter.ToObject(val)
	}
	if ref == "location" {
		hasChange = true
		locations := strings.Split(botMsg.Text, ";")
		inputfilter.Locations = append(inputfilter.Locations, locations...)
	}
	if ref == "type" {
		hasChange = true
		inputfilter.Type = strings.TrimSpace(botMsg.Text)
	}
	if ref == "currency" {
		hasChange = true
		inputfilter.Currency = strings.TrimSpace(botMsg.Text)
	}
	if ref == "salary" {
		hasChange = true
		salaryArr := strings.Split(botMsg.Text, "-")
		if len(salaryArr) >= 1 {
			inputfilter.SalaryFrom = salaryArr[0]
			if len(salaryArr) == 2 {
				inputfilter.SalaryTo = salaryArr[1]
			}
		}
	}

	if ref == "inline-arg" {
		_, params := textparse.QueryParse(botMsg.Text)
		for _, items := range *params {
			for key, item := range items {
				if key == "location[]" {
					inputfilter.Locations = append(inputfilter.Locations, item)
				}
				if key == "type" {
					inputfilter.Type = item
				}
				if key == "currency" {
					inputfilter.Currency = item
				}
				if key == "salaryFrom" {
					inputfilter.SalaryFrom = item
				}
				if key == "salaryTo" {
					inputfilter.SalaryTo = item
				}
			}
		}
	}

	if val, err := inputfilter.ToJSON(); err == nil && hasChange {

		err := redis.Set(ctx, keyID, val, 8*time.Hour).Err()
		if err != nil {
			log.Println(err.Error())
		} else {
			var bot = setup.BOT
			msgConfig := tgbotapi.NewMessage(botMsg.Chat.ID, "")
			msgConfig.Text = "set next filter or close ?"
			msgConfig.ReplyMarkup = filterKeyboard

			if _, err := bot.Send(msgConfig); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
