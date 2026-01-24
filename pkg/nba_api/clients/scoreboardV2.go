package clients

import (
	"fmt"
)

const scoreBoardV2UrlFormat = "scoreboardv2?GameDate=%s&LeagueID=00&DayOffset=0"

func (apiClient *NbaApiClient) FetchScoreBoardV2(date string) error {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(scoreBoardV2UrlFormat, date)

	err := sendCommonGetRequest(apiClient, url)
	return err
}
