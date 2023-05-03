package telegram

import (
	"encoding/json"
	"fmt"
)

func GetWebhookInfoFunc(config TGApiConfig) (string, error) {

	bot, err := newBotWithConfig(config)
	if err != nil {
		return "", fmt.Errorf("error to create bot: %v", err)
	}

	webhookInfo, err := bot.GetWebhookInfo()
	if err != nil {
		return "", fmt.Errorf("error to get webhook info: %v", err)
	}

	webhookJson, _ := json.MarshalIndent(webhookInfo, "", "  ")
	return fmt.Sprintf("%s\n", webhookJson), nil
}

func SetWebhookInfoFunc(config TGApiConfig, request SetWebhookConfig) (string, error) {

	bot, err := newBotWithConfig(config)
	if err != nil {
		return "", fmt.Errorf("error to create bot: %v", err)
	}

	resp, err := bot.MakeRequest(request.Method(), request.Params())
	if err != nil {
		return "", fmt.Errorf("error to set webhook: %v", err)
	}

	webhookJson, _ := json.MarshalIndent(resp, "", "  ")
	return fmt.Sprintf("%s\n", webhookJson), nil
}
