package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const boxScoreV3UrlFormat = "boxscoretraditionalv3?GameID=%s&StartPeriod=0&EndPeriod=0&StartRange=0&EndRange=0&RangeType=0"

func (apiClient *NbaApiClient) FetchBoxScoreTraditionalV3JSON(gameID string) (json.RawMessage, error) {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(boxScoreV3UrlFormat, gameID)

	req, err := buildCommonGetRequest(url)

	response, err := apiClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nba stats HTTP %d: %s", response.StatusCode, string(body))
	}

	return body, nil
}
