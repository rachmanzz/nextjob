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
	//var redis = setup.REDIS

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// var numericKeyboard = tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("filter"),
	// 		tgbotapi.NewKeyboardButton("2"),
	// 		tgbotapi.NewKeyboardButton("3"),
	// 	),
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("4"),
	// 		tgbotapi.NewKeyboardButton("5"),
	// 		tgbotapi.NewKeyboardButton("6"),
	// 	),
	// )

	// var filterButton = tgbotapi.NewInlineKeyboardMarkup(
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardRowButtonData("location", "--filter-location"),
	// 		tgbotapi.NewInlineKeyboardButtonData("salary start", "--filter-salary-from"),
	// 		tgbotapi.NewInlineKeyboardButtonData("salary until", "--filter-salary-end"),
	// 	),
	// 	tgbotapi.NewInlineKeyboardRow(
	// 		tgbotapi.NewInlineKeyboardButtonData("currency", "--filter-currency"),
	// 		tgbotapi.NewInlineKeyboardButtonData("experience", "--filter-experience"),
	// 		tgbotapi.NewInlineKeyboardButtonData("next", "--filter-next"),
	// 	),
	// )

	debounced := debounce.New(2000 * time.Millisecond)

	for update := range updates {
		if update.Message != nil { // If we got a message
			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			if update.Message.IsCommand() { // ignore any non-command Messages
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

				// article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", update.InlineQuery.Query)
				// article.Description = update.InlineQuery.Query

				// inlineConf := tgbotapi.InlineConfig{
				// 	InlineQueryID: update.InlineQuery.ID,
				// 	IsPersonal:    true,
				// 	CacheTime:     0,
				// 	Results:       []interface{}{article},
				// }

				// if _, err := bot.Request(inlineConf); err != nil {
				// 	log.Println(err)
				// }
			}
		}
	}
}
