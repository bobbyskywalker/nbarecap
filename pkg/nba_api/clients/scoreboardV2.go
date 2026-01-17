package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const scoreBoardV2UrlFormat = "scoreboardv2?GameDate=%s&LeagueID=00&DayOffset=0"

func (apiClient *NbaApiClient) FetchScoreboardV2(date string) (json.RawMessage, error) {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(scoreBoardV2UrlFormat, date)
	req, err := buildCommonGetRequest(url)
	if err != nil {
		return nil, err
	}

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
