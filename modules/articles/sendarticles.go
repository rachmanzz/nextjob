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

		content := "<b><u>" + title + "</u></b>\n"

		if itemType, ok := item["type"].(string); ok && itemType != "" {
			content += "<b>type</b> " + itemType + "\n"
		}
		if itemArrangement, ok := item["arrangement"].(string); ok && itemArrangement != "" {
			content += "<b>arrangement</b> " + itemArrangement + "\n"
		}

		if itemExperience, ok := item["experience"].(string); ok && itemExperience != "" {
			content += "<b>experience</b> " + itemExperience + "\n"
		}
		if salaryItem, ok := item["salary"].(map[string]interface{}); ok {
			if salary, ok := salaryItem["range"].(string); ok && salary != "" {
				content += fmt.Sprintf("<b>salary</b> %s\n", salary)
			}
		}

		if itemCurrency, ok := item["currency"].(string); ok && itemCurrency != "" {
			content += "<b>currency</b> " + itemCurrency + "\n"
		}

		if itemLocation, ok := item["location"].(string); ok && itemLocation != "" {
			content += "<b>location</b> " + itemLocation + "\n"
		}

		if stacks, ok := item["stacks"].([]interface{}); ok {
			stacklist := ""
			for _, stackItem := range stacks {
				if stackData, ok := stackItem.(map[string]interface{}); ok {
					if stackName, ok := stackData["name"].(string); ok {
						if stacklist != "" {
							stacklist += ", " + stackName
						} else {
							stacklist = stackName
						}
					}
				}
			}

			if stacklist != "" {
				content += "\n<u>Tech Stack:</u>\n" + stacklist + "\n\n"
			}
		}

		fullContent := fmt.Sprintf("%s \n%s", content, contentURL)

		article := tgbotapi.NewInlineQueryResultArticleHTML(articleID, title, fullContent)
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
