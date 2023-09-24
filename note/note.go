package note

import "time"

type Note struct {
	ChatId       string    `json:"chat_id"`
	Text         string    `json:"text"`
	ReceivedTime time.Time `json:"received_time"`
}
