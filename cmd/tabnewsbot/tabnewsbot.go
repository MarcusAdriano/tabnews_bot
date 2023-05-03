package tabnewsbot

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/cmd/telegram"
	"github.com/spf13/cobra"
)

var isLambdaMode = false
var telegramBotCmd = &cobra.Command{
	Use:   "telegram",
	Short: "Run telegram bot",
	Run: func(cmd *cobra.Command, args []string) {

		token := os.Getenv("TELEGRAM_BOT_TOKEN")
		url := os.Getenv("TELEGRAM_API_URL")
		if len(url) == 0 {
			url = tgbotapi.APIEndpoint
		}

		httpClient := http.Client{}

		config := telegram.TGApiConfig{
			URL:        url,
			Token:      token,
			HttpClient: httpClient,
		}

		if isLambdaMode {
			telegram.RunTGBotLambdaModeFunc(config)
		} else {
			telegram.RunTGBotPollingModeFunc(config)
		}
	},
}

var rootCmd = &cobra.Command{
	Use:   "tabnewsbot",
	Short: "Run TabNews chat bot",
}

func initialize() {
	telegramBotCmd.Flags().BoolVarP(&isLambdaMode, "lambda", "l", false, "Run in lambda mode")

	rootCmd.AddCommand(telegramBotCmd)
}

func Execute() {
	initialize()

	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error to send telegram message: %v\n", err)
	}
}
