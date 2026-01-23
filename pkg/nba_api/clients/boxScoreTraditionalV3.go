package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const boxScoreV3UrlFormat = "boxscoretraditionalv3?GameID=%s&StartPeriod=0&EndPeriod=0&StartRange=0&EndRange=0&RangeType=0"

func (apiClient *NbaApiClient) FetchBoxScoreTraditionalV3Json(gameID string) (json.RawMessage, error) {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(boxScoreV3UrlFormat, gameID)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setCommonRequestHeaders(req)
	return getCommonJsonResponse(apiClient, req)
}
