package handler

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/articles"
	"github.com/rachmanzz/nextjob/modules/httprequest"
	"github.com/rachmanzz/nextjob/modules/textparse"
	"github.com/rachmanzz/nextjob/setup"
	"github.com/rachmanzz/nextjob/types"
)

func HandleInlieQuery(inline *tgbotapi.InlineQuery) error {
	var redis = setup.REDIS
	var query = inline.Query
	if query == "" {
		return nil
	}
	positions, params := textparse.QueryParse(query)

	keyID := fmt.Sprintf("filter_%d", inline.From.ID)

	var filterParams []map[string]string = []map[string]string{}

	if val, err := redis.Get(ctx, keyID).Result(); err == nil && (params == nil || len(*params) == 0) {
		inputfilter := types.FilterRefObject{}
		inputfilter.ToObject(val)

		if inputfilter.Locations != nil && len(inputfilter.Locations) >= 1 {
			for _, val := range inputfilter.Locations {
				filterParams = append(filterParams, map[string]string{"location[]": val})
			}
		}
		if inputfilter.Currency != "" {
			filterParams = append(filterParams, map[string]string{"currency": inputfilter.Currency})
		}
		if inputfilter.SalaryFrom != "" {
			filterParams = append(filterParams, map[string]string{"salaryFrom": inputfilter.SalaryFrom})
		}
		if inputfilter.SalaryTo != "" {
			filterParams = append(filterParams, map[string]string{"salaryTo": inputfilter.SalaryTo})
		}
		if inputfilter.Type != "" {
			filterParams = append(filterParams, map[string]string{"typ": inputfilter.Type})
		}

	} else {
		if params != nil && len(*params) >= 1 {
			filterParams = append(filterParams, *params...)
		}
	}

	res, err := httprequest.JobRequest(*positions, filterParams)

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
