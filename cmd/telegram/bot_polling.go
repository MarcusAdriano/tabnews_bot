package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/telegram"
)

func RunTGBotPollingMode(config TGApiConfig) {

	bot, err = tgbotapi.NewBotAPIWithClient(config.Token, config.URL, &config.HttpClient)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

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
