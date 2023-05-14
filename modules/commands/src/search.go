package src

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/commands/options"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

var ctx = context.Background()

func SearchCommand(msg *tgbotapi.MessageConfig, opts ...any) {
	var redis = setup.REDIS
	var bot = setup.BOT
	opt := options.GetOpt(opts...)
	str := fmt.Sprintf("Type job title that you looking for %d", opt.ChatID)
	msg.Text = str

	var command string = "search"
	var typedID = "typing_" + strconv.Itoa(int(opt.ChatID))

	if message, err := bot.Send(msg); err == nil {
		searchData := types.InputMessageObject{
			ChatID:          msg.ChatID,
			MessageID:       message.MessageID,
			Command:         &command,
			ChannelUsername: msg.ChannelUsername,
		}
		redis.Set(ctx, typedID, *searchData.ToJSON(), time.Second*20)
		go timeoutTyping(searchData)
	}
}

func timeoutTyping(searchData types.InputMessageObject) {
	time.Sleep(time.Second * 15)
	var bot = setup.BOT
	_, err := bot.Request(tgbotapi.DeleteMessageConfig{
		ChannelUsername: searchData.ChannelUsername,
		ChatID:          searchData.ChatID,
		MessageID:       searchData.MessageID,
	})
	if err != nil {
		log.Println(err.Error())
	}

}
