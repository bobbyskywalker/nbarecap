package scores

import (
	"errors"
	"fmt"
	"nbarecap/internal/models"
	"nbarecap/internal/utils"
	"sort"
	"strings"

	"github.com/ronaudinho/nag"
)

const (
	TITLE = "Today's NBA games\n\n"

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
		id := utils.AnyAsString(r[H_GAME_ID])
		g := &models.GameInfo{
			GameID:   id,
			Status:   utils.AnyAsString(r[H_STATUS_TEXT]),
			GameCode: utils.AnyAsString(r[H_GAMECODE]),
			Arena:    utils.AnyAsString(r[H_ARENA_NAME]),
			NatTV:    utils.AnyAsString(r[H_NATL_TV]),
			HomeTV:   utils.AnyAsString(r[H_HOME_TV]),
			AwayTV:   utils.AnyAsString(r[H_AWAY_TV]),
			HomeID:   utils.AnyAsString(r[H_HOME_TEAM_ID]),
			AwayID:   utils.AnyAsString(r[H_VIS_TEAM_ID]),
		}
		g.AwayAbbr, g.HomeAbbr = fallbackAbbrsFromGameCode(g.GameCode)
		g.SortKey = buildSortKey(g.Status, g.GameID)

		gameMap[id] = g
	}

	if lsAny, ok := responseMap[RS_LINESCORE]; ok {
		if lsRows, ok := lsAny.([]map[string]any); ok {
			for _, r := range lsRows {
				gameID := utils.AnyAsString(r[H_GAME_ID])
				teamID := utils.AnyAsString(r[H_TEAM_ID])

				g := gameMap[gameID]
				if g == nil {
					continue
				}

				info := models.TeamInfo{
					Abbr:   utils.AnyAsString(r[H_TEAM_ABBR]),
					Record: utils.AnyAsString(r[H_TEAM_WINS_LOSSES]),
					Pts:    utils.AnyAsIntPtr(r[H_PTS]),
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

func FormatGamesForDay(res nag.Response) (string, error) {
	responseMap := nag.Map(res)

	ghAny, ok := responseMap[RS_GAMEHEADER]
	if !ok {
		return "", errors.New(HEADER_MISSING_GAMEHEADER)
	}
	ghRows, ok := ghAny.([]map[string]any)
	if !ok {
		return "", errors.New(HEADER_BAD_GAMEHEADER)
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

	return b.String(), nil
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
