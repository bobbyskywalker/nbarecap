package nba

import (
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/mappers"
	"nbarecap/pkg/nba_api/models"
)

func GetStandingsForSeason(season string) ([]*models.Standing, error) {
	client := clients.NewNbaApiClient()
	err := client.FetchStandingsV3(season)
	if err != nil {
		return nil, fmt.Errorf("error fetching standings for season %s: %w", season, err)
	}

	standingsList, err := mappers.BuildStandingsList(client.Response.ResultSets)
	if err != nil {
		return nil, err
	}

	return standingsList, nil
}
