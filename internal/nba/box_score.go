package nba

import (
	"errors"
	"fmt"
	"log"
	"nbarecap/pkg/nba_api/clients"
	"os"
)

func FetchBoxScoreForGame(gameId string) (*string, error) {
	client := clients.NewNbaApiClient()
	bx, err := client.FetchBoxScoreTraditionalV3JSON(gameId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error fetching box score for gameId %s: %v", gameId, err))
	}

	file, _ := os.Create("box.json")
	_, err = file.Write(bx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return nil, nil
}
