package telegram

import (
	"encoding/json"
	"fmt"
)

func GetWebhookInfo(config TGApiConfig) (string, error) {

	bot, err := newBotWithConfig(config)
	if err != nil {
		return "", fmt.Errorf("error to create bot: %v", err)
	}

	webhookInfo, err := bot.GetWebhookInfo()
	return jsonIdent(webhookInfo, err)
}

func SetWebhookInfo(config TGApiConfig, request SetWebhookConfig) (string, error) {

	bot, err := newBotWithConfig(config)
	if err != nil {
		return "", fmt.Errorf("error to create bot: %v", err)
	}

	resp, err := bot.MakeRequest(request.Method(), request.Params())
	return jsonIdent(resp, err)
}

func DeleteWebhook(config TGApiConfig, request DeleteWebhookConfig) (string, error) {

	bot, err := newBotWithConfig(config)
	if err != nil {
		return "", fmt.Errorf("error to create bot: %v", err)
	}

	resp, err := bot.MakeRequest(request.Method(), request.Params())
	return jsonIdent(resp, err)
}

func jsonIdent(result interface{}, err error) (string, error) {

	if err != nil {
		return "", fmt.Errorf("error to get result info: %v", err)
	}

	webhookJson, _ := json.MarshalIndent(result, "", "  ")
	return fmt.Sprintf("%s\n", webhookJson), nil
}
