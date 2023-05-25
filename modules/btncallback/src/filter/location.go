package filter

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

var ctx = context.Background()

func FilterLocation(msg tgbotapi.Message) {
	var bot = setup.BOT
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")

	msgConfig.Text = "Please type location. If there more than one location, split by semicolon (;)"

	nextMess, err := bot.Send(msgConfig)
	if err != nil {
		log.Println(err.Error())
	}

	var redis = setup.REDIS

	command := "/filter"
	refData := "location"

	input := types.InputMessageObject{
		ChatID:    nextMess.Chat.ID,
		MessageID: nextMess.MessageID,
		Command:   &command,
		RefData:   &refData,
	}

	keyID := fmt.Sprintf("typing_%d", input.ChatID)

	if val, err := input.ToJSON(); err == nil {
		err := redis.Set(ctx, keyID, val, 15*time.Second).Err()
		if err != nil {
			log.Println(err.Error())
		}

		go func() {
			time.Sleep(time.Second * 12)
			if _, err := redis.Get(ctx, keyID).Result(); err == nil {
				bot.Request(tgbotapi.DeleteMessageConfig{
					MessageID: nextMess.MessageID,
					ChatID:    nextMess.Chat.ID,
				})
			}
		}()
	}

}
