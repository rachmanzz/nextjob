package handler

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/input"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

var ctx = context.Background()

func HandleInput(botMsg *tgbotapi.Message) {
	var redis = setup.REDIS

	keyID := fmt.Sprintf("typing_%d", botMsg.Chat.ID)
	if val, err := redis.Get(ctx, keyID).Result(); err == nil && botMsg.Text != "" {
		inputData := types.InputMessageObject{}
		if err := inputData.ToObject(val); err == nil {
			for _, v := range input.RegisteredInputCommand {
				if v.CommandName == *inputData.Command {
					v.CommandAction(*inputData.RefData, botMsg)
					redis.Del(ctx, keyID)
				}
			}
		}

	}
}
