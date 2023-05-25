package types

import (
	"encoding/json"
)

type FilterRefObject struct {
	Locations  []string `json:"locations"`
	Type       string   `json:"type"`
	Currency   string   `json:"currency"`
	SalaryFrom string   `json:"salary_from"`
	SalaryTo   string   `json:"salary_to"`
}

func (input FilterRefObject) ToJSON() (string, error) {
	var jData, err = json.Marshal(input)
	if err != nil {
		return "", err
	}

	jsonStr := string(jData)

	return jsonStr, nil
}

func (input *FilterRefObject) ToObject(val string) error {
	err := json.Unmarshal([]byte(val), input)
	if err != nil {
		return err
	}
	return nil
}
