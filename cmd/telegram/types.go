package telegram

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGApiConfig struct {
	URL        string
	Token      string
	HttpClient http.Client
}

type SetWebhookConfig struct {
	URL                string   `json:"url"`
	Certificate        string   `json:"certificate,omitempty"`
	MaxConnections     int      `json:"max_connections,omitempty"`
	AllowedUpdates     []string `json:"allowed_updates,omitempty"`
	DropPendingUpdates bool     `json:"drop_pending_updates,omitempty"`
	SecretToken        string   `json:"secret_token,omitempty"`
	IpAddress          string   `json:"ip_address,omitempty"`
}

func (s SetWebhookConfig) Params() tgbotapi.Params {
	params := make(tgbotapi.Params)

	params.AddNonEmpty("url", s.URL)
	params.AddNonEmpty("certificate", s.Certificate)
	params.AddNonZero("max_connections", s.MaxConnections)
	params.AddInterface("allowed_updates", s.AllowedUpdates)
	params.AddBool("drop_pending_updates", s.DropPendingUpdates)
	params.AddNonEmpty("secret_token", s.SecretToken)
	params.AddNonEmpty("ip_address", s.IpAddress)

	return params
}

func (s SetWebhookConfig) Method() string {
	return "setWebhook"
}
