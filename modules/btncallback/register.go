package btncallback

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rachmanzz/nextjob/modules/btncallback/src/filter"
)

type RegisteredQueryType struct {
	Query          string
	CallbackAction func(msg tgbotapi.Message)
}

var RegisteredQuery []RegisteredQueryType

func init() {
	RegisteredQuery = []RegisteredQueryType{
		// filter query
		{
			Query:          "--filter-location",
			CallbackAction: filter.FilterLocation,
		},
		{
			Query:          "--filter-type",
			CallbackAction: filter.FilterType,
		},
		{
			Query:          "--filter-currency",
			CallbackAction: filter.FilterCurrency,
		},
		{
			Query:          "--filter-salary",
			CallbackAction: filter.FilterSalary,
		},
		{
			Query:          "--filter-inline-arg",
			CallbackAction: filter.FilterInlineArgument,
		},
		{
			Query:          "--filter-clear",
			CallbackAction: filter.FilterClear,
		},
		{
			Query:          "--filter-open",
			CallbackAction: filter.FilterOpen,
		},
		{
			Query:          "--filter-close",
			CallbackAction: filter.FilterBtnFinish,
		},
		{
			Query:          "--filter-data",
			CallbackAction: filter.FilterData,
		},
		// end filter query
	}
}
