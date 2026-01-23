package nba

import (
	"errors"
	"fmt"
	"nbarecap/pkg/nba_api/clients"
	"nbarecap/pkg/nba_api/mappers"
	"nbarecap/pkg/nba_api/models"
	"sort"
	"strings"
	"time"
)

const (
	title                       = " NBA games\n\n"
	dateFormat                  = "2006-01-02"
	scoresFormat                = " — %d-%d"
	gameInfoFormat              = "%s — %s%s @ %s%s%s\n"
	headerMissingGameheader     = "No GameHeader in response.\n"
	headerBadGameheader         = "GameHeader has unexpected type.\n"
	rsGameHeader                = "GameHeader"
	nationalBroadcastInfoFormat = "National: %s"
	awayBroadcastInfoFormat     = "Away: %s"
	homeBroadcastInfoFormat     = "Home: %s"
)

func formatGamesForDay(date *string, responseMap map[string]any) ([]models.GameInfoFormatted, error) {
	var result []models.GameInfoFormatted

	ghAny, ok := responseMap[rsGameHeader]
	if !ok {
		return nil, errors.New(headerMissingGameheader)
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return nil, errors.New(headerBadGameheader)
	}

	gameMap := mappers.BuildGameMap(responseMap, ghRows)

	list := make([]*models.GameInfo, 0, len(gameMap))
	for _, game := range gameMap {
		list = append(list, game)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].SortKey < list[j].SortKey
	})

	result = append(result, models.NewGameInfoFormatted("", *date+title))

	for _, game := range list {
		sb := strings.Builder{}

		awayAbbr := game.Away.Abbr
		homeAbbr := game.Home.Abbr
		if awayAbbr == "" {
			awayAbbr = game.AwayAbbr
		}
		if homeAbbr == "" {
			homeAbbr = game.HomeAbbr
		}

		awayRec := formatRecord(game.Away.Record)
		homeRec := formatRecord(game.Home.Record)

		score := ""
		if game.Away.Pts != nil && game.Home.Pts != nil {
			score = fmt.Sprintf(scoresFormat, *game.Away.Pts, *game.Home.Pts)
		}

		_, _ = fmt.Fprintf(&sb, gameInfoFormat,
			strings.TrimSpace(game.Status),
			awayAbbr, awayRec,
			homeAbbr, homeRec,
			score,
		)

		tv := formatTVLine(game)
		if tv != "" {
			_, _ = fmt.Fprintf(&sb, "  %s\n", tv)
		}
		if strings.TrimSpace(game.Arena) != "" {
			_, _ = fmt.Fprintf(&sb, "  %s\n", strings.TrimSpace(game.Arena))
		}

		result = append(result, models.NewGameInfoFormatted(game.GameID, sb.String()))
	}

	return result, nil
}

func formatRecord(rec string) string {
	rec = strings.TrimSpace(rec)
	if rec == "" {
		return ""
	}
	return " (" + rec + ")"
}

func formatTVLine(g *models.GameInfo) string {
	parts := make([]string, 0, 3)

	if s := strings.TrimSpace(g.NatTV); s != "" {
		parts = append(parts, fmt.Sprintf(nationalBroadcastInfoFormat, s))
	}
	if s := strings.TrimSpace(g.HomeTV); s != "" {
		parts = append(parts, fmt.Sprintf(homeBroadcastInfoFormat, s))
	}
	if s := strings.TrimSpace(g.AwayTV); s != "" {
		parts = append(parts, fmt.Sprintf(awayBroadcastInfoFormat, s))
	}
	return strings.Join(parts, " • ")
}

func GetAllGamesForDate(date *time.Time) ([]models.GameInfoFormatted, error) {
	dateStr := ""
	if date != nil {
		dateStr = (*date).Format(dateFormat)
	} else {
		dateStr = time.Now().Format(dateFormat)
	}

	client := clients.NewNbaApiClient()
	rsMap, err := client.FetchScoreboardV2(dateStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error fetching games for date %s: %v", dateStr, err))
	}
	return formatGamesForDay(&dateStr, rsMap)
}
