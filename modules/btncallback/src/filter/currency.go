package filter

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

func FilterCurrency(msg tgbotapi.Message) {

	var bot = setup.BOT
	msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "")

	msgConfig.Text = "Please type currency as usd, idr, sgd, etc..."

	nextMess, err := bot.Send(msgConfig)
	if err != nil {
		log.Println(err.Error())
	}

	var redis = setup.REDIS

	command := "/filter"
	refData := "currency"

	input := types.InputMessageObject{
		ChatID:    nextMess.Chat.ID,
		MessageID: nextMess.MessageID,
		Command:   &command,
		RefData:   &refData,
	}

	keyID := fmt.Sprintf("typing_%d", input.ChatID)

	if val, err := input.ToJSON(); err == nil {
		err := redis.Set(ctx, keyID, val, 30*time.Second).Err()
		if err != nil {
			log.Println(err.Error())
		}

		go func() {
			time.Sleep(time.Second * 25)
			if _, err := redis.Get(ctx, keyID).Result(); err == nil {
				bot.Request(tgbotapi.DeleteMessageConfig{
					MessageID: nextMess.MessageID,
					ChatID:    nextMess.Chat.ID,
				})
			}
		}()
	}

}
