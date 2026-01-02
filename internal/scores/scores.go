package scores

import (
	"fmt"
	"nbarecap/internal/models"
	"sort"
	"strings"

	"github.com/ronaudinho/nag"
)

const (
	/* Parsing & Formatting */
	ABBR_LEN      = 3
	ABBR_PAIR_LEN = ABBR_LEN * 2

	SORT_KEY_SEP = "|"
	SPLIT_SEP    = "/"

	SCORES_FORMAT    = " — %d-%d"
	MAIN_INFO_FORMAT = "%s — %s%s @ %s%s%s\n"

	HEADER_MISSING_GAMEHEADER = "No GameHeader in response.\n"
	HEADER_BAD_GAMEHEADER     = "GameHeader has unexpected type.\n"

	/* Result set names */
	RS_GAMEHEADER = "GameHeader"
	RS_LINESCORE  = "LineScore"

	/* Column / header names */
	H_GAME_ID      = "GAME_ID"
	H_STATUS_TEXT  = "GAME_STATUS_TEXT"
	H_GAMECODE     = "GAMECODE"
	H_ARENA_NAME   = "ARENA_NAME"
	H_NATL_TV      = "NATL_TV_BROADCASTER_ABBREVIATION"
	H_HOME_TV      = "HOME_TV_BROADCASTER_ABBREVIATION"
	H_AWAY_TV      = "AWAY_TV_BROADCASTER_ABBREVIATION"
	H_HOME_TEAM_ID = "HOME_TEAM_ID"
	H_VIS_TEAM_ID  = "VISITOR_TEAM_ID"

	H_TEAM_ID          = "TEAM_ID"
	H_TEAM_ABBR        = "TEAM_ABBREVIATION"
	H_TEAM_WINS_LOSSES = "TEAM_WINS_LOSSES"
	H_PTS              = "PTS"
)

func asString(v any) string {
	if v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprintf("%v", t)
	}
}

func asIntPtr(v any) *int {
	if v == nil {
		return nil
	}
	if f, ok := v.(float64); ok {
		n := int(f)
		return &n
	}
	if n, ok := v.(int); ok {
		return &n
	}
	return nil
}

func fallbackAbbrsFromGameCode(gameCode string) (away, home string) {
	tokens := strings.Split(gameCode, SPLIT_SEP)
	if len(tokens) != 2 {
		return "", ""
	}

	teams := tokens[1]
	if len(teams) == ABBR_PAIR_LEN {
		return teams[:ABBR_LEN], teams[ABBR_LEN:]
	}
	return "", ""
}

func buildSortKey(status, gameID string) string {
	return status + SORT_KEY_SEP + gameID
}

func buildGameMap(responseMap map[string]any, ghRows []map[string]any) map[string]*models.GameInfo {
	gameMap := make(map[string]*models.GameInfo, len(ghRows))

	for _, r := range ghRows {
		id := asString(r[H_GAME_ID])
		g := &models.GameInfo{
			GameID:   id,
			Status:   asString(r[H_STATUS_TEXT]),
			GameCode: asString(r[H_GAMECODE]),
			Arena:    asString(r[H_ARENA_NAME]),
			NatTV:    asString(r[H_NATL_TV]),
			HomeTV:   asString(r[H_HOME_TV]),
			AwayTV:   asString(r[H_AWAY_TV]),
			HomeID:   asString(r[H_HOME_TEAM_ID]),
			AwayID:   asString(r[H_VIS_TEAM_ID]),
		}
		g.AwayAbbr, g.HomeAbbr = fallbackAbbrsFromGameCode(g.GameCode)
		g.SortKey = buildSortKey(g.Status, g.GameID)

		gameMap[id] = g
	}

	if lsAny, ok := responseMap[RS_LINESCORE]; ok {
		if lsRows, ok := lsAny.([]map[string]any); ok {
			for _, r := range lsRows {
				gameID := asString(r[H_GAME_ID])
				teamID := asString(r[H_TEAM_ID])

				g := gameMap[gameID]
				if g == nil {
					continue
				}

				info := models.TeamInfo{
					Abbr:   asString(r[H_TEAM_ABBR]),
					Record: asString(r[H_TEAM_WINS_LOSSES]),
					Pts:    asIntPtr(r[H_PTS]),
				}

				switch teamID {
				case g.HomeID:
					g.Home = info
				case g.AwayID:
					g.Away = info
				}
			}
		}
	}

	return gameMap
}

func FormatTodaysGames(res nag.Response) string {
	responseMap := nag.Map(res)

	ghAny, ok := responseMap[RS_GAMEHEADER]
	if !ok {
		return HEADER_MISSING_GAMEHEADER
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return HEADER_BAD_GAMEHEADER
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
	b.WriteString("Today's NBA games\n\n") // const

	for _, g := range list {
		awayAbbr := g.Away.Abbr
		homeAbbr := g.Home.Abbr
		if awayAbbr == "" {
			awayAbbr = g.AwayAbbr
		}
		if homeAbbr == "" {
			homeAbbr = g.HomeAbbr
		}

		awayRec := strings.TrimSpace(g.Away.Record)
		homeRec := strings.TrimSpace(g.Home.Record)
		if awayRec != "" {
			awayRec = " (" + awayRec + ")"
		}
		if homeRec != "" {
			homeRec = " (" + homeRec + ")"
		}

		score := ""
		if g.Away.Pts != nil && g.Home.Pts != nil {
			score = fmt.Sprintf(SCORES_FORMAT, *g.Away.Pts, *g.Home.Pts)
		}

		fmt.Fprintf(&b, MAIN_INFO_FORMAT, g.Status, awayAbbr, awayRec, homeAbbr, homeRec, score)
		fmt.Fprintf(&b, "  %s\n\n", g.Arena)
	}

	return b.String()
}

func GetAllGamesForDate(date *string) error {
	sb := nag.NewScoreBoardV2()
	if date != nil {
		sb.GameDate = *date
	}

	if err := sb.Get(); err != nil {
		return err
	}

	fmt.Println(FormatTodaysGames(*sb.Response))
	return nil
}
