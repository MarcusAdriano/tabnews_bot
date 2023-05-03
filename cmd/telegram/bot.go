package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/telegram"
)

var bot *tgbotapi.BotAPI
var err error

func RunBotPollingModeFunc(config TGApiConfig) {

	bot, err = tgbotapi.NewBotAPIWithClient(config.Token, config.URL, &config.HttpClient)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			message := telegram.TabNewsTgBotUpdate{
				Update: update,
				Sender: sender,
			}

			go telegram.ReceiveMessage(message)
		}
	}
}

func sender(messages []tgbotapi.MessageConfig) {
	for _, msg := range messages {
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Cannot send message %v, error: %v\n", msg, err)
		}
	}
}
