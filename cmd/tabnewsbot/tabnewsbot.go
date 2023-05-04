package tabnewsbot

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marcusadriano/tabnews_bot/cmd/telegram"
	"github.com/spf13/cobra"
)

func startTelegramBot(cmd *cobra.Command, args []string) {

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

	if isLocalMode {
		telegram.RunTGBotPollingMode(config)
	} else {
		telegram.RunTGBotLambdaMode(config)
	}
}

var isLocalMode = false
var telegramBotCmd = &cobra.Command{
	Use:   "telegram",
	Short: "Run telegram bot",
	Run:   startTelegramBot,
}

var rootCmd = &cobra.Command{
	Use:   "tabnewsbot",
	Short: "Run TabNews chat bot",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: startTelegramBot,
}

func initialize() {
	telegramBotCmd.Flags().BoolVarP(&isLocalMode, "local", "l", false, "Run in local mode")
	rootCmd.Flags().BoolVarP(&isLocalMode, "local", "l", false, "Run in local mode")
	rootCmd.AddCommand(telegramBotCmd)
}

func Execute() {
	initialize()

	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error to send telegram message: %v\n", err)
	}
}
