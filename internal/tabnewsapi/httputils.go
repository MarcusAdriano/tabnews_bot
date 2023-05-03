package tabnewsapi

import (
	"fmt"
	"net/url"
)

func addNonEmptyParam(params url.Values, key string, value string) {
	if value != "" {
		params.Add(key, value)
	}
}

func addNonZeroParam(params url.Values, key string, value int) {
	if value != 0 {
		params.Add(key, fmt.Sprintf("%d", value))
	}
}

func buildUrlWithParams(url *url.URL, params url.Values) string {
	url.RawQuery = params.Encode()
	return url.String()
}
