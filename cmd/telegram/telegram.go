package telegram

import (
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

		config := TGApiConfig{
			URL:        telegramApiUrl,
			Token:      token,
			HttpClient: httpClient,
		}

		result, err := GetWebhookInfoFunc(config)
		if err != nil {
			finalizeWithError(err)
		}

		fmt.Println(result)
	},
}

var setWebhookConfig SetWebhookConfig

var setWebhookCmd = &cobra.Command{
	Use:   "set",
	Short: "Set webhook details",
	Run: func(cmd *cobra.Command, args []string) {

		config := TGApiConfig{
			URL:        telegramApiUrl,
			Token:      token,
			HttpClient: httpClient,
		}

		result, err := SetWebhookInfoFunc(config, setWebhookConfig)
		if err != nil {
			finalizeWithError(err)
		}

		fmt.Println(result)
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

func newBotWithConfig(config TGApiConfig) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPIWithClient(config.Token, config.URL, &config.HttpClient)
}

func initialize() {
	webhookCmd.PersistentFlags().StringVarP(&telegramApiUrl, "telegramApiUrl", "u", tgbotapi.APIEndpoint, "Telegram API URL")
	webhookCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Telegram bot token")
	webhookCmd.MarkPersistentFlagRequired("token")

	setWebhookCmd.PersistentFlags().StringVarP(&setWebhookConfig.URL, "webhook", "w", "", "Webhook URL")
	setWebhookCmd.PersistentFlags().IntVarP(&setWebhookConfig.MaxConnections, "max_connections", "m", 40, "Max connections")
	setWebhookCmd.PersistentFlags().StringArrayVarP(&setWebhookConfig.AllowedUpdates, "allowed_updates", "a", []string{}, "Allowed updates")
	setWebhookCmd.PersistentFlags().BoolVarP(&setWebhookConfig.DropPendingUpdates, "drop_pending_updates", "d", false, "Drop pending updates")
	setWebhookCmd.PersistentFlags().StringVarP(&setWebhookConfig.SecretToken, "secret_token", "s", "", "Secret token")
	setWebhookCmd.MarkPersistentFlagRequired("webhook")

	webhookCmd.AddCommand(getWebhookInfoCmd, setWebhookCmd)

	rootCmd.AddCommand(webhookCmd)
}

func finalizeWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
