package bot

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"noteavard/database"
	"noteavard/note"
	"strconv"
	"time"
)

const botToken = "6563971667:AAF4dmpkmPYFCVlTXyWN_38Pem-glv0WnJI"
const baseUrl = "https://api.telegram.org/bot"

func sendMessage(chatId int, text string) error {
	result, err := http.PostForm(baseUrl+botToken+"/sendMessage", url.Values{
		"chat_id": {strconv.Itoa(chatId)},
		"text":    {text},
	})

	if err != nil {
		log.Println(err)
		return err
	}

	defer result.Body.Close()

	return nil
}

func SaveNote(chatId int, message string) {
	note := note.Note{
		Text:         message,
		ChatId:       strconv.Itoa(chatId),
		ReceivedTime: time.Now(),
	}

	database.DbInstance.Create(&note)

	sendMessage(chatId, message)
}

func SendNotes(chatId int, noteDate time.Time) {

	var notes []note.Note

	database.DbInstance.Where("received_time between ? and ?", fmt.Sprintf("%d-%d-%d", noteDate.Year(), noteDate.Month(), noteDate.Day()), fmt.Sprintf("%d-%d-%d", noteDate.Year(), noteDate.Month(), noteDate.Day()+1)).Find(&notes)

	fmt.Println(notes)

	messageText := createNoteMessageText(notes)
	sendMessage(chatId, messageText)
}

func createNoteMessageText(notes []note.Note) string {

	message := "یادداشت های امروز 📆:"

	for _, note := range notes {
		message = message + fmt.Sprintf("%s•\n", note.Text)
	}

	return message
}
