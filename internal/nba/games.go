package nba

import (
	"errors"
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/mappers"
	"nbarecap/pkg/nba_api/models"
	"sort"
	"time"
)

const (
	dateFormat              = "2006-01-02"
	headerMissingGameheader = "No GameHeader in response.\n"
	headerBadGameheader     = "GameHeader has unexpected type.\n"
	rsGameHeader            = "GameHeader"
)

func GetAllGamesForDate(date *time.Time) ([]*models.GameInfo, error) {
	dateStr := ""
	if date != nil {
		dateStr = (*date).Format(dateFormat)
	} else {
		dateStr = time.Now().Format(dateFormat)
	}

	client := clients.NewNbaApiClient()
	err := client.FetchScoreBoardV2(dateStr)
	if err != nil {
		return nil, fmt.Errorf("error fetching scoreboard for %s: %w", dateStr, err)
	}

	ghAny, ok := client.Response.ResultSets[rsGameHeader]
	if !ok {
		return nil, errors.New(headerMissingGameheader)
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return nil, errors.New(headerBadGameheader)
	}

	gameMap := mappers.BuildGameMap(client.Response.ResultSets, ghRows)

	list := make([]*models.GameInfo, 0, len(gameMap))
	for _, game := range gameMap {
		list = append(list, game)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].SortKey < list[j].SortKey
	})

	return list, nil
}
