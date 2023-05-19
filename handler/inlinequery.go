package handler

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/articles"
	"github.com/rachmanzz/nextjob/modules/httprequest"
	"github.com/rachmanzz/nextjob/modules/textparse"
)

func HandleInlieQuery(inline *tgbotapi.InlineQuery) error {
	var query = inline.Query
	if query == "" {
		return nil
	}
	positions, params := textparse.QueryParse(query)

	res, err := httprequest.JobRequest(*positions, *params)

	if err != nil {
		log.Println("error request to showwcase network")
		return err
	}
	defer res.Body.Close()

	var data []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Println("error decode to json")
		return err
	}

	return articles.SendArticles(inline.ID, data)

}
