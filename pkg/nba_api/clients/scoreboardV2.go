package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const scoreBoardV2UrlFormat = "scoreboardv2?date=%s"

func (apiClient *NbaApiClient) FetchScoreboardV2(date string) (json.RawMessage, error) {
	// to fix
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(scoreBoardV2UrlFormat, date)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setRequestHeaders(req)

	response, err := apiClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nba stats HTTP %d", response.StatusCode)
	}

	return body, nil
}
