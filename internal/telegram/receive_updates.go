package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/tabnewsapi"
)

const (
	TabNewsBaseApiUrl = "https://www.tabnews.com.br"
	DefaultPageSize   = 10
	DefaultPage       = 0
)

var tabNewsApi = tabnewsapi.NewTabNewsAPI(TabNewsBaseApiUrl)

type TGBotSender = func(bot tgbotapi.Chattable)

type TabNewsTgBotUpdate struct {
	Update tgbotapi.Update
	Sender TGBotSender
}

func ReceiveMessage(message TabNewsTgBotUpdate) {

	update := message.Update.Message
	if update.IsCommand() {
		receiveCommand(message)
		return
	} else {
		receiveHelp(message)
	}
}

func receiveCommand(message TabNewsTgBotUpdate) {

	update := message.Update
	switch update.Message.Command() {
	case "start":
		receiveStart(message)
	case "relevant":
		receiveRelevant(message)
	case "old":
		receiveOld(message)
	case "new":
		receiveNews(message)
	default:
		receiveUnknownCommand(message)
	}
}

func receiveStart(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Olá, seja bem-vindo ao TabNews Bot! Utilize os comandos e seja feliz ;D")

	message.Sender(msg)
}

func receiveRelevant(message TabNewsTgBotUpdate) {

	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyRelevant,
	}
	newInlineButtonResponse(message, contentConfig)
}

func receiveNews(message TabNewsTgBotUpdate) {
	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyNew,
	}
	newInlineButtonResponse(message, contentConfig)
}

func receiveOld(message TabNewsTgBotUpdate) {
	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyOld,
	}
	newInlineButtonResponse(message, contentConfig)
}

func newInlineButtonResponse(message TabNewsTgBotUpdate, config tabnewsapi.ContentsConfig) {

	update := message.Update
	contents, err := tabNewsApi.Contents(config)
	if err != nil {
		log.Fatalf("Error calling TabNews API: %v\n", err)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Desculpe, não consegui acessar o TabNews agora :(")
		message.Sender(msg)
		return
	}

	var inlineKeyboardRows [][]tgbotapi.InlineKeyboardButton

	for _, content := range contents {
		contentUrl := content.Link(TabNewsBaseApiUrl)
		button := tgbotapi.NewInlineKeyboardButtonURL(content.Title, contentUrl)
		row := tgbotapi.NewInlineKeyboardRow(button)
		inlineKeyboardRows = append(inlineKeyboardRows, row)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Segue o posts:")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboardRows...)

	message.Sender(msg)
}

func receiveHelp(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Utilize um dos comandos abaixo.\n\nComandos:\n/relevant - Notícias mais relevantes\n/old - Notícias mais antigas\n/new - Notícias mais recentes")

	message.Sender(msg)
}

func receiveUnknownCommand(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Comando desconhecido")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender(msg)
}
