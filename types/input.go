package types

import (
	"encoding/json"
)

type InputMessageObject struct {
	ChatID          int64   `json:"chat_id"`
	MessageID       int     `json:"message_id"`
	ChannelUsername string  `json:"channel_username"`
	Command         *string `json:"command"`
}

func (input InputMessageObject) ToJSON() *string {
	var jData, err = json.Marshal(input)
	if err != nil {
		return nil
	}

	jsonStr := string(jData)

	return &jsonStr
}

func (input *InputMessageObject) ToObject(val string) error {
	err := json.Unmarshal([]byte(val), input)
	if err != nil {
		return err
	}
	return nil
}
