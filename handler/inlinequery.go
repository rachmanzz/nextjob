package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

func HandleInlieQuery(inline *tgbotapi.InlineQuery) error {
	var bot = setup.BOT
	var query = inline.Query

	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://cache.showwcase.com/jobs", nil)
	if err != nil {
		log.Println("make request error")
	}

	request.Header.Set("x-api-key", "d40f0824bb5f30856db3a9e5285d74c1345745b71a09acfe69")

	q := request.URL.Query()
	q.Add("position[]", query)
	q.Add("limit", "6")
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)

	if err != nil {
		log.Println("make on execute network")
	}
	defer response.Body.Close()

	var data []map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("error decode to json")
		return err
	}

	var articles []interface{}

	for _, val := range data {
		var title = val["title"].(string)
		var desc = val["description"].(string)
		if len(desc) > 200 {
			desc = string(desc[0 : 200-1])
		}
		id := val["id"].(float64)
		slug := val["slug"].(string)
		artID := fmt.Sprintf("id-%d", int(id))
		company := val["company"].(map[string]interface{})
		contentURL := fmt.Sprintf("https://www.showwcase.com/job/%d-%s", int(id), slug)
		article := tgbotapi.NewInlineQueryResultArticle(artID, title, contentURL)
		article.Description = desc
		if logo, ok := company["logo"].(string); ok {
			article.ThumbURL = logo
		}
		if applyURL, ok := val["applyUrl"].(string); ok {
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("APPLY", applyURL),
				),
			)
			article.ReplyMarkup = &keyboard
		}

		articles = append(articles, article)
	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       articles,
	}

	if _, err := bot.Request(inlineConf); err != nil {
		log.Println(err)
	}

	return nil

}
