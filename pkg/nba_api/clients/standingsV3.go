package clients

import "fmt"

const (
	standingsV3UrlFormat = "leaguestandingsv3?LeagueID=00&Season=%s&SeasonType=%s"
	regularSeasonParam   = "Regular%20Season"
)

func (apiClient *NbaApiClient) FetchStandingsV3(season string) error {
	url := apiClient.baseUrl + apiClient.statsSuffix + fmt.Sprintf(standingsV3UrlFormat, season, regularSeasonParam)
	err := sendCommonGetRequest(apiClient, url)
	return err
}
