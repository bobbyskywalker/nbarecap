package nba

import (
	"errors"
	"fmt"
	"nbarecap/internal/models"
	"sort"
	"strings"
	"time"

	"github.com/ronaudinho/nag"
)

func formatGamesForDay(date *string, response nag.Response) ([]models.GameInfoFormatted, error) {
	responseMap := nag.Map(response)
	var result []models.GameInfoFormatted

	ghAny, ok := responseMap[RsGameheader]
	if !ok {
		return nil, errors.New(HeaderMissingGameheader)
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return nil, errors.New(HeaderBadGameheader)
	}

	gameMap := buildGameMap(responseMap, ghRows)

	list := make([]*models.GameInfo, 0, len(gameMap))
	for _, game := range gameMap {
		list = append(list, game)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].SortKey < list[j].SortKey
	})

	result = append(result, models.NewGameInfoFormatted("", *date+Title))

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
			score = fmt.Sprintf(ScoresFormat, *game.Away.Pts, *game.Home.Pts)
		}

		_, _ = fmt.Fprintf(&sb, GameInfoFormat,
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
		parts = append(parts, fmt.Sprintf(NationalBroadcastInfoFormat, s))
	}
	if s := strings.TrimSpace(g.HomeTV); s != "" {
		parts = append(parts, fmt.Sprintf(HomeBroadcastInfoFormat, s))
	}
	if s := strings.TrimSpace(g.AwayTV); s != "" {
		parts = append(parts, fmt.Sprintf(AwayBroadcastInfoFormat, s))
	}
	return strings.Join(parts, " • ")
}

func GetAllGamesForDate(date *time.Time) ([]models.GameInfoFormatted, error) {
	dateStr := ""
	sb := nag.NewScoreBoardV2()

	if date != nil {
		dateStr = (*date).Format(DateFormat)
		sb.GameDate = dateStr
	}
	if err := sb.Get(); err != nil {
		return nil, errors.New(fmt.Sprintf("Error fetching games: %v", err))
	}
	return formatGamesForDay(&dateStr, *sb.Response)
}
