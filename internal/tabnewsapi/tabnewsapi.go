package tabnewsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TabNewsAPI interface {
	Contents(config ContentsConfig) ([]Content, error)
}

type tabNewsAPI struct {
	baseUrl string
	client  *http.Client
}

func NewTabNewsAPI(baseUrl string) TabNewsAPI {

	client := &http.Client{}

	return tabNewsAPI{
		baseUrl: baseUrl,
		client:  client,
	}
}

func (t tabNewsAPI) Contents(config ContentsConfig) ([]Content, error) {

	urlBase, err := url.Parse(t.baseUrl + "/api/v1/contents")
	if err != nil {
		return []Content{}, err
	}

	method := "GET"

	params := createUrlParams(config)
	tabNewsUrl := buildUrlWithParams(urlBase, params)

	fmt.Println("Calling TabNews API with URL:", tabNewsUrl)

	req, err := http.NewRequest(method, tabNewsUrl, nil)
	req.Header.Set("user-agent", "Telegram BOT (tabnews_bot; v0.1) ")

	if err != nil {
		return []Content{}, err
	}
	res, err := t.client.Do(req)
	if err != nil {
		return []Content{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []Content{}, err
	}

	var contents []Content
	err = json.Unmarshal(body, &contents)
	return contents, err
}

func createUrlParams(config ContentsConfig) url.Values {
	params := url.Values{}

	addNonZeroParam(params, "page", config.Page)
	addNonZeroParam(params, "per_page", config.PerPage)
	addNonEmptyParam(params, "strategy", config.Strategy)

	return params
}
