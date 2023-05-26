package main

import (
	"log"
	"time"

	"github.com/bep/debounce"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/handler"
	"github.com/rachmanzz/nextjob/setup"
)

func main() {
	var bot = setup.BOT

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	debounced := debounce.New(2000 * time.Millisecond)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				handler.HandleCommand(update.Message)
			} else {
				handler.HandleInput(update.Message)
			}
		} else {
			if update.CallbackQuery != nil {
				handler.HandleCallbackQuery(update)
			}

			if update.InlineQuery != nil {
				debounced(func() { handler.HandleInlieQuery(update.InlineQuery) })
			}
		}
	}
}
