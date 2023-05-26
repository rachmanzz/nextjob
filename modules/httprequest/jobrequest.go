package httprequest

import (
	"net/http"
	"strings"

	"github.com/rachmanzz/nextjob/setup"
)

func JobRequest(positions []string, params []map[string]string) (*http.Response, error) {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "https://cache.showwcase.com/jobs", nil)
	if err != nil {
		return nil, err
	}

	if setup.VarArgData.ShowCaseAPI != nil {
		request.Header.Set("x-api-key", *setup.VarArgData.ShowCaseAPI)
	}

	q := request.URL.Query()
	for _, position := range positions {
		q.Add("position[]", strings.TrimSpace(position))
	}

	for _, param := range params {
		for key, val := range param {
			q.Add(key, val)
		}
	}

	q.Add("limit", "6")
	request.URL.RawQuery = q.Encode()

	return client.Do(request)
}
