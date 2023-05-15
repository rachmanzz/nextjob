package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands"
	"github.com/rachmanzz/nextjob/setup"
)

func HandleCommand(botMsg *tgbotapi.Message) {
	var bot = setup.BOT
	msg := tgbotapi.NewMessage(botMsg.Chat.ID, "")
	for _, v := range commands.RegisteredCommand {
		if v.CommandName == botMsg.Command() {
			var name string = botMsg.Chat.UserName
			if botMsg.Chat.FirstName != "" {
				name = botMsg.Chat.FirstName
				if botMsg.Chat.LastName != "" {
					name += " " + botMsg.Chat.LastName
				}
			}

			bot.Request(tgbotapi.DeleteMessageConfig{
				MessageID: botMsg.MessageID,
				ChatID:    botMsg.Chat.ID,
			})

			v.CommandAction(&msg, name, botMsg.Chat.ID, botMsg.MessageID, botMsg.Chat.UserName)

		}
	}
}
