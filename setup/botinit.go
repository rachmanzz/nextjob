package setup

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var BOT *tgbotapi.BotAPI
var REDIS *redis.Client

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("SECRET"))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	BOT = bot

	REDIS = rdb
}
