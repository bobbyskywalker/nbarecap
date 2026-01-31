package nba

import (
	"encoding/json"
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/models"
)

func GetPlayByPlayForGame(gameId string) (*models.PlayByPlayV3, error) {
	client := clients.NewNbaApiClient()
	err := client.FetchPlayByPlayV3FullGame(gameId)
	if err != nil {
		return nil, fmt.Errorf("error fetching playByPlay for gameID %s: %w", gameId, err)
	}
	result := models.PlayByPlayV3{}
	err = json.Unmarshal(client.Response.Json, &result)
	return &result, err
}
