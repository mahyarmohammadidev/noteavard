package main

import (
	"encoding/json"
	"log"
	"net/http"
	"noteavard/bot"
	"noteavard/database"
	"time"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	Id int `json:"id"`
}

func main() {

	database.ConnectToSqlite()
	database.Migrate()

	http.ListenAndServe(":8080", http.HandlerFunc(Handler))
}

func Handler(response http.ResponseWriter, request *http.Request) {

	var incommingMessage, err = parseAsTelegramRequest(request)

	log.Println(incommingMessage.Message.Text)

	if err != nil {
		return
	}

	switch incommingMessage.Message.Text {
	case "/todaynotes":
		bot.SendNotes(incommingMessage.Message.Chat.Id, time.Now())
	default:
		bot.SaveNote(incommingMessage.Message.Chat.Id, incommingMessage.Message.Text)
	}
}

func parseAsTelegramRequest(req *http.Request) (*Update, error) {
	var update Update

	err := json.NewDecoder(req.Body).Decode(&update)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &update, nil
}
