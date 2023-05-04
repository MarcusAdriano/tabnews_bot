package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI
var err error

func sender(msg tgbotapi.Chattable) {
	res, err := bot.Send(msg)
	if err != nil {
		log.Printf("Cannot send message %v, error: %v\n", msg, err)
	} else {
		log.Printf("Message was sent to user: %d\n", res.Chat.ID)
	}
}
