package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ReceiveMessage(update tgbotapi.Update) {
	fmt.Println(update.Message)
}
