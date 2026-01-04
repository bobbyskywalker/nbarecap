package recaps

import (
	"errors"
	"fmt"
	"nbarecap/internal/models"
	"sort"
	"strings"
	"time"

	"github.com/ronaudinho/nag"
)

func FormatGamesForDay(date *string, response nag.Response) ([]string, error) {
	responseMap := nag.Map(response)
	var result []string

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

	result = append(result, *date+Title)
	sb := strings.Builder{}

	for _, game := range list {
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

		_, _ = fmt.Fprintf(&sb, "%s — %s%s @ %s%s%s\n",
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

		result = append(result, sb.String())
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
		parts = append(parts, "National: "+s)
	}
	if s := strings.TrimSpace(g.HomeTV); s != "" {
		parts = append(parts, "Home: "+s)
	}
	if s := strings.TrimSpace(g.AwayTV); s != "" {
		parts = append(parts, "Away: "+s)
	}
	return strings.Join(parts, " • ")
}

func GetAllGamesForDate(date *time.Time) ([]string, error) {
	dateStr := ""
	sb := nag.NewScoreBoardV2()

	if date != nil {
		dateStr = (*date).Format("2006-01-02")
		sb.GameDate = dateStr
	}
	if err := sb.Get(); err != nil {
		return nil, nil
	}
	return FormatGamesForDay(&dateStr, *sb.Response)
}
