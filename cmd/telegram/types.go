package telegram

import "net/http"

type TGApiConfig struct {
	URL        string
	Token      string
	HttpClient http.Client
}
