package clients

import "fmt"

const playByPlayV3Format = "playbyplayv3?EndPeriod=%s&GameID=%s&StartPeriod=%s"

func (apiClient *NbaApiClient) FetchPlayByPlayV3(gameID string, endPeriod string, startPeriod string) error {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(playByPlayV3Format, endPeriod, gameID, startPeriod)
	err := sendCommonGetRequest(apiClient, url)
	return err
}

func (apiClient *NbaApiClient) FetchPlayByPlayV3FullGame(gameID string) error {
	return apiClient.FetchPlayByPlayV3(gameID, "10", "1")
}
