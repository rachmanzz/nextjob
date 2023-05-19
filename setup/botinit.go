package setup

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var BOT *tgbotapi.BotAPI

func init() {
	godotenv.Load("local.env")

	err := PopulateVarArg()

	if err != nil {
		panic(err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(*VarArgData.BotAPI)
	if err != nil {
		panic(err)
	}

	if !VarArgData.Production {
		bot.Debug = true
	}

	BOT = bot

	RunRedis()
}
