package filter

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

func FilterClear(msg tgbotapi.Message) {
	var redis = setup.REDIS
	keyID := fmt.Sprintf("filter_%d", msg.Chat.ID)

	var bot = setup.BOT
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")

	err := redis.Del(ctx, keyID).Err()
	if err != nil {
		msgConfig.Text = "failed to clear filter."
	} else {
		msgConfig.Text = "cleared...."
	}

	if _, err := bot.Send(msgConfig); err != nil {
		log.Println(err.Error())
	}
}
