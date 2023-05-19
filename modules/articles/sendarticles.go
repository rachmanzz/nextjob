package articles

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/setup"
)

func SendArticles(queryID string, data []map[string]interface{}) error {
	var bot = setup.BOT
	var articleList []interface{}

	for _, item := range data {
		title := item["title"].(string)
		desc := item["description"].(string)
		if len(desc) > 300 {
			desc = string(desc[0 : 300-1])
		}

		itemID := item["id"].(float64)
		slug := item["slug"].(string)
		articleID := fmt.Sprintf("id-%d", int(itemID))
		contentURL := fmt.Sprintf("https://www.showwcase.com/job/%d-%s", int(itemID), slug)

		content := title + "\n"

		if itemType, ok := item["type"].(string); ok && itemType != "" {
			content += "type : " + itemType + "\n"
		}
		if itemArrangement, ok := item["arrangement"].(string); ok && itemArrangement != "" {
			content += "arrangement : " + itemArrangement + "\n"
		}

		if itemExperience, ok := item["experience"].(string); ok && itemExperience != "" {
			content += "experience : " + itemExperience + "\n"
		}

		salaryFrom := item["salaryFrom"].(float64)
		salaryTo := item["salaryTo"].(float64)
		content += fmt.Sprintf("salary : %.f-%.f\n", salaryFrom, salaryTo)

		if itemCurrency, ok := item["currency"].(string); ok && itemCurrency != "" {
			content += "currency : " + itemCurrency + "\n"
		}

		if itemLocation, ok := item["location"].(string); ok && itemLocation != "" {
			content += "location : " + itemLocation + "\n"
		}

		fullContent := fmt.Sprintf("%s \n%s", contentURL, content)

		article := tgbotapi.NewInlineQueryResultArticle(articleID, title, fullContent)
		article.Description = desc

		if company, ok := item["company"].(map[string]interface{}); ok {
			if logo, ok := company["logo"].(string); ok {
				article.ThumbURL = logo
			}
			if applyURL, ok := item["applyUrl"].(string); ok && strings.TrimSpace(applyURL) != "" {
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("APPLY", applyURL),
					),
				)
				article.ReplyMarkup = &keyboard
			}
		}

		articleList = append(articleList, article)

	}

	inlineConf := tgbotapi.InlineConfig{
		InlineQueryID: queryID,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       articleList,
	}

	if _, err := bot.Request(inlineConf); err != nil {
		return err
	}

	return nil

}
