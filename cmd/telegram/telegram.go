package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/cobra"
)

var httpClient http.Client = http.Client{}
var token string
var telegramApiUrl string

var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Manage telegram webhook details",
	Long:  `You can get, delete and set telegram's webhook.`,
}

var getWebhookInfoCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current webhook details",
	Run: func(cmd *cobra.Command, args []string) {

		bot, err := tgbotapi.NewBotAPIWithClient(token, telegramApiUrl, &httpClient)
		if err != nil {
			finalizeWithMessageAndError("Error to create Telegram Bot API client: ", err)
		}

		webhookInfo, err := bot.GetWebhookInfo()
		if err != nil {
			finalizeWithMessageAndError("Error to get webhook info: ", err)
		}

		json, _ := json.Marshal(webhookInfo)
		fmt.Printf("%s\n", json)
	},
}

var rootCmd = &cobra.Command{
	Use:   "telegram",
	Short: "Manage telegram webhook details",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("telegram called")
	},
}

func Execute() {

	initialize()

	if err := rootCmd.Execute(); err != nil {
		finalizeWithError(err)
	}
}

func initialize() {
	webhookCmd.PersistentFlags().StringVarP(&telegramApiUrl, "telegramApiUrl", "u", tgbotapi.APIEndpoint, "Telegram API URL")
	webhookCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Telegram bot token")
	webhookCmd.MarkPersistentFlagRequired("token")

	webhookCmd.AddCommand(getWebhookInfoCmd)

	rootCmd.AddCommand(webhookCmd)
}

func finalizeWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func finalizeWithMessageAndError(message string, err error) {
	fmt.Println(message, err)
	os.Exit(1)
}
