package clients

import (
	"fmt"
)

const boxScoreV3UrlFormat = "boxscoretraditionalv3?GameID=%s&StartPeriod=0&EndPeriod=0&StartRange=0&EndRange=0&RangeType=0"

func (apiClient *NbaApiClient) FetchBoxScoreTraditionalV3(gameID string) error {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(boxScoreV3UrlFormat, gameID)
	err := sendCommonGetRequest(apiClient, url)
	return err
}
