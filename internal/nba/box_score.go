package nba

import (
	"encoding/json"
	"errors"
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/models"
)

func GetBoxScoreForGame(gameId string) (*models.BoxScoreTraditionalV3, error) {
	client := clients.NewNbaApiClient()
	err := client.FetchBoxScoreTraditionalV3(gameId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error fetching box score for game %s: %v", gameId, err))
	}
	var boxScore models.BoxScoreTraditionalV3Response
	err = json.Unmarshal(client.Response.Json, &boxScore)
	if err != nil {
		return nil, err
	}
	return &boxScore.BoxScoreTraditional, nil
}
