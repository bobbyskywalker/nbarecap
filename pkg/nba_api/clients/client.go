package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"nbarecap/pkg/nba_api/mappers"
	"net/http"
	"time"
)

const defaultTimeout = 20 * time.Second

type NbaApiResponse struct {
	ResultSets map[string]any
	Json       json.RawMessage
}

func newNbaApiResponse() *NbaApiResponse {
	return &NbaApiResponse{
		ResultSets: make(map[string]any),
		Json:       nil,
	}
}

type NbaApiClient struct {
	baseUrl     string
	statsSuffix string
	httpClient  *http.Client
	Response    *NbaApiResponse
}

func NewNbaApiClient() *NbaApiClient {
	return &NbaApiClient{
		baseUrl:     "https://stats.nba.com/",
		statsSuffix: "stats/",
		httpClient:  &http.Client{Timeout: defaultTimeout},
		Response:    newNbaApiResponse(),
	}
}

func setCommonRequestHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Referer", "https://www.nba.com/")
	req.Header.Set("Origin", "https://www.nba.com")
}

func sendCommonGetRequest(apiClient *NbaApiClient, url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	setCommonRequestHeaders(req)

	response, err := apiClient.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("nba stats HTTP %d", response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	apiClient.Response.Json = body
	apiClient.Response.ResultSets, err = mappers.MapResultSetsToResponseMap(body)

	return err
}
