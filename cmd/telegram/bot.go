package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/pkg/telegram"
)

func RunBotPollingModeFunc(config TGApiConfig) {

	bot, err := tgbotapi.NewBotAPIWithClient(config.Token, config.URL, &config.HttpClient)
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
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			go telegram.ReceiveMessage(update)
		}
	}
}
