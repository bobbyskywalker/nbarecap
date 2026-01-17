package clients

import (
	"context"
	"fmt"
	"net/http"
)

const scoreBoardV2UrlFormat = "scoreboardv2?GameDate=%s&LeagueID=00&DayOffset=0"

func (apiClient *NbaApiClient) FetchScoreboardV2(date string) (map[string]any, error) {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(scoreBoardV2UrlFormat, date)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setCommonRequestHeaders(req)
	return getResultSetsFromJson(apiClient, req)
}
