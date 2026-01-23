package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"nbarecap/pkg/nba_api/mappers"
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

func setCommonRequestHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")
}

// TODO: refactor for two ways of fetching

func getCommonJsonResponse(apiClient *NbaApiClient, request *http.Request) (json.RawMessage, error) {
	response, err := apiClient.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("nba stats HTTP %d", response.StatusCode))
	}
	return body, nil
}

func getResultSetsFromJson(apiClient *NbaApiClient, request *http.Request) (map[string]any, error) {
	response, err := apiClient.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("nba stats HTTP %d", response.StatusCode))
	}

	return mappers.MapResultSetsToResponseMap(body)
}
