package types

import (
	"encoding/json"
)

type InputMessageObject struct {
	ChatID    int64   `json:"chat_id"`
	MessageID int     `json:"message_id"`
	Command   *string `json:"command"`
	RefData   *string `json:"ref_data"`
}

func (input InputMessageObject) ToJSON() (string, error) {
	var jData, err = json.Marshal(input)
	if err != nil {
		return "", err
	}

	jsonStr := string(jData)

	return jsonStr, nil
}

func (input *InputMessageObject) ToObject(val string) error {
	err := json.Unmarshal([]byte(val), input)
	if err != nil {
		return err
	}
	return nil
}
