package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/internal/telegram"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	telegramSecret := request.Headers["X-Telegram-Bot-Api-Secret-Token"]
	if telegramSecret != os.Getenv("TELEGRAM_BOT_API_SECRET_TOKEN") {
		return events.APIGatewayProxyResponse{Body: "Unauthorized", StatusCode: 401}, nil
	}

	var update tgbotapi.Update
	err := json.Unmarshal([]byte(request.Body), &update)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "Bad Request", StatusCode: 400}, nil
	}

	message := telegram.TabNewsTgBotUpdate{
		Update: update,
		Sender: sender,
	}

	telegram.ReceiveMessage(message)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func RunTGBotLambdaMode(config TGApiConfig) {

	bot, err = tgbotapi.NewBotAPIWithClient(config.Token, config.URL, &config.HttpClient)
	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(handleRequest)
}
