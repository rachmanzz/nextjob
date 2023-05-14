package handler

import (
	"context"
	"log"

	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

var ctx = context.Background()

func HandleInput(botMsg *tgbotapi.Message) {
	var bot = setup.BOT
	var redis = setup.REDIS
	msg := tgbotapi.NewMessage(botMsg.Chat.ID, "")
	if botMsg.Text != "" {
		if val, err := redis.Get(ctx, "typing_"+strconv.Itoa(int(botMsg.Chat.ID))).Result(); err == nil {
			commandSearch := types.InputMessageObject{}
			if commandSearch.ToObject(val); err != nil {
				log.Panic(err)
			}
			if *commandSearch.Command == "search" {
				msg.Text = "in search"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
