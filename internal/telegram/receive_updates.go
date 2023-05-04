package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/tabnewsapi"
)

const (
	TabNewsBaseApiUrl = "https://www.tabnews.com.br"
	DefaultPageSize   = 10
	DefaultPage       = 0
)

const (
	fire = `🔥`
	new  = `🆕`
	old  = `👴`
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
		handleCommand(message)
		return
	} else {
		handleHelp(message)
	}
}

func handleCommand(message TabNewsTgBotUpdate) {

	update := message.Update
	switch update.Message.Command() {
	case "start":
		handleStartCmd(message)
	case "relevant":
		handleRelevantCmd(message)
	case "old":
		handleOldCmd(message)
	case "new":
		handleNewCmd(message)
	default:
		handleUnknownCmd(message)
	}
}

func handleStartCmd(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Olá, seja bem-vindo ao TabNews Bot! Utilize os comandos e seja feliz ;D")

	message.Sender(msg)
}

func handleRelevantCmd(message TabNewsTgBotUpdate) {

	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyRelevant,
	}
	newInlineButtonResponse(message, fmt.Sprintf("%s Relevantes: %s", fire, fire), contentConfig)
}

func handleNewCmd(message TabNewsTgBotUpdate) {
	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyNew,
	}
	newInlineButtonResponse(message, fmt.Sprintf("%s Novos posts: %s", new, new), contentConfig)
}

func handleOldCmd(message TabNewsTgBotUpdate) {
	contentConfig := tabnewsapi.ContentsConfig{
		Page:     DefaultPage,
		PerPage:  DefaultPageSize,
		Strategy: tabnewsapi.StrategyOld,
	}
	newInlineButtonResponse(message, fmt.Sprintf("%s Posts antigos: %s", old, old), contentConfig)
}

func newInlineButtonResponse(message TabNewsTgBotUpdate, label string, config tabnewsapi.ContentsConfig) {

	update := message.Update
	contents, err := tabNewsApi.Contents(config)
	if err != nil || len(contents) == 0 {
		log.Printf("Error calling TabNews API: %v\n", err)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Desculpe, não consegui acessar o TabNews agora :(")
		message.Sender(msg)
		return
	}

	var inlineKeyboardRows [][]tgbotapi.InlineKeyboardButton

	for _, content := range contents {
		contentUrl := content.Link(TabNewsBaseApiUrl)
		webAppInfo := tgbotapi.WebAppInfo{
			URL: contentUrl,
		}

		button := tgbotapi.NewInlineKeyboardButtonWebApp(content.Title, webAppInfo)
		row := tgbotapi.NewInlineKeyboardRow(button)
		inlineKeyboardRows = append(inlineKeyboardRows, row)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, label)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboardRows...)

	message.Sender(msg)
}

func handleHelp(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Utilize um dos comandos abaixo.\n\nComandos:\n/relevant - Notícias mais relevantes\n/old - Notícias mais antigas\n/new - Notícias mais recentes")

	message.Sender(msg)
}

func handleUnknownCmd(message TabNewsTgBotUpdate) {
	update := message.Update
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Comando desconhecido")
	msg.ReplyToMessageID = update.Message.MessageID

	message.Sender(msg)
}
