package clients

import (
	"context"
	"net/http"
	"time"
)

const defaultTimeout = 20 * time.Second

type NbaApiClient struct {
	baseUrl     string
	statsSuffix string
	httpClient  *http.Client
}

func NewNbaApiClient() *NbaApiClient {
	return &NbaApiClient{
		baseUrl:     "https://stats.nba.com/",
		statsSuffix: "stats/",
		httpClient:  &http.Client{Timeout: defaultTimeout},
	}
}

func buildCommonGetRequest(url string) (*http.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setRequestHeaders(req)
	return req.WithContext(ctx), nil
}

func setRequestHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")
}
