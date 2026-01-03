package recaps

import (
	"errors"
	"fmt"
	"nbarecap/internal/models"
	"sort"
	"strings"

	"github.com/ronaudinho/nag"
)

func FormatGamesForDay(res nag.Response) (string, error) {
	responseMap := nag.Map(res)

	ghAny, ok := responseMap[RsGameheader]
	if !ok {
		return "", errors.New(HeaderMissingGameheader)
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return "", errors.New(HeaderBadGameheader)
	}

	gameMap := buildGameMap(responseMap, ghRows)

	list := make([]*models.GameInfo, 0, len(gameMap))
	for _, g := range gameMap {
		list = append(list, g)
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].SortKey < list[j].SortKey
	})

	var b strings.Builder
	b.WriteString(TITLE)

	for _, g := range list {
		awayAbbr := g.Away.Abbr
		homeAbbr := g.Home.Abbr
		if awayAbbr == "" {
			awayAbbr = g.AwayAbbr
		}
		if homeAbbr == "" {
			homeAbbr = g.HomeAbbr
		}

		awayRec := formatRecord(g.Away.Record)
		homeRec := formatRecord(g.Home.Record)

		score := ""
		if g.Away.Pts != nil && g.Home.Pts != nil {
			score = fmt.Sprintf(ScoresFormat, *g.Away.Pts, *g.Home.Pts)
		}

		_, _ = fmt.Fprintf(&b, "%s — %s%s @ %s%s%s\n",
			strings.TrimSpace(g.Status),
			awayAbbr, awayRec,
			homeAbbr, homeRec,
			score,
		)

		tv := formatTVLine(g)
		if tv != "" {
			_, _ = fmt.Fprintf(&b, "  %s\n", tv)
		}
		if strings.TrimSpace(g.Arena) != "" {
			_, _ = fmt.Fprintf(&b, "  %s\n", strings.TrimSpace(g.Arena))
		}

		b.WriteString("\n")
	}

	return b.String(), nil
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

func GetAllGamesForDate(date *string) (string, error) {
	sb := nag.NewScoreBoardV2()
	if date != nil {
		sb.GameDate = *date
	}
	if err := sb.Get(); err != nil {
		return "", nil
	}
	return FormatGamesForDay(*sb.Response)
}
