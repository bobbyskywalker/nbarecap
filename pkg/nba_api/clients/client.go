package clients

import (
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

func setRequestHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("x-nba-stats-origin", "stats")
	req.Header.Set("x-nba-stats-token", "true")
}
