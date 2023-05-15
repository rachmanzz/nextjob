package setup

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var BOT *tgbotapi.BotAPI

func init() {
	if os.Getenv("LOCAL_ENV") == "yes" {
		err := godotenv.Load("local.env")
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("no local env")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("SECRET"))
	if err != nil {
		panic(err)
	}

	BOT = bot

	RunRedis()
}
