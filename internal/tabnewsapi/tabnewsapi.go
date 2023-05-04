package tabnewsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/marcusadriano/tabnews_bot/internal/cache"
)

const (
	cacheTtl = time.Minute * 10
)

type TabNewsAPI interface {
	Contents(config ContentsConfig) ([]Content, error)
}

type tabNewsAPI struct {
	baseUrl string
	client  *http.Client
	cache   *cache.Cache
}

func NewTabNewsAPI(baseUrl string) TabNewsAPI {

	client := &http.Client{}
	inMemoryCache := cache.NewCache()

	return tabNewsAPI{
		baseUrl: baseUrl,
		client:  client,
		cache:   inMemoryCache,
	}
}

func (t tabNewsAPI) Contents(config ContentsConfig) ([]Content, error) {

	return t.contentsFromCache(config)
}

func (t tabNewsAPI) contentsFromCache(config ContentsConfig) ([]Content, error) {

	key := fmt.Sprintf("%s-%d-%d", config.Strategy, config.Page, config.PerPage)

	if v, ok := t.cache.Get(key); ok {
		log.Printf("Cache hit for key: %s\n", key)
		return v.([]Content), nil
	}

	contents, err := t.contentsFromSource(config)
	if err == nil && len(contents) > 0 {
		t.cache.Set(key, contents, cacheTtl)
	}
	return contents, err
}

func (t tabNewsAPI) contentsFromSource(config ContentsConfig) ([]Content, error) {

	urlBase, err := url.Parse(t.baseUrl + "/api/v1/contents")
	if err != nil {
		return []Content{}, err
	}

	method := "GET"

	params := createUrlParams(config)
	tabNewsUrl := buildUrlWithParams(urlBase, params)

	log.Printf("Calling TabNews API with URL: %s\n", tabNewsUrl)

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
