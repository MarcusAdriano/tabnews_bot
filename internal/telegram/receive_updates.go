package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/tabnewsapi"
)

var tabNewsApi = tabnewsapi.NewTabNewsAPI("https://www.tabnews.com.br")

type TGBotSender = func(bot []tgbotapi.MessageConfig)

type TabNewsTgBotUpdate struct {
	Update tgbotapi.Update
	Sender TGBotSender
}

func ReceiveMessage(message TabNewsTgBotUpdate) {

	update := message.Update.Message
	if update.IsCommand() {
		receiveCommand(message)
		return
	}

	//contentConfig := tabnewsapi.ContentsConfig{
	//	Page:     0,
	//	PerPage:  10,
	//	Strategy: tabnewsapi.StrategyRelevant,
	//}
	//contents, err := tabNewsApi.Contents(contentConfig)
	//
	//if err != nil {
	//	fmt.Println("Error calling TabNews API:", err)
	//}

}

func receiveCommand(message TabNewsTgBotUpdate) {

	update := message.Update
	switch update.Message.Command() {
	case "start":
		receiveStart(message)
	case "relevant":
		receiveRelevant(message)
	case "old":
		receiveHelp(message)
	case "new":
		receiveNews(message)
	default:
		receiveUnknownCommand(message)
	}
}

func receiveStart(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Olá, seja bem-vindo ao TabNews Bot! Utilize os comandos e seja feliz ;D")

	message.Sender([]tgbotapi.MessageConfig{msg})
}

func receiveRelevant(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Relevant")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender([]tgbotapi.MessageConfig{msg})
}

func receiveHelp(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Help")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender([]tgbotapi.MessageConfig{msg})
}

func receiveNews(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "News")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender([]tgbotapi.MessageConfig{msg})
}

func receiveUnknownCommand(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Comando desconhecido")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender([]tgbotapi.MessageConfig{msg})
}
