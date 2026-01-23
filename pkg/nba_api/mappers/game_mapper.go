package mappers

import (
	"nbarecap/internal/utils"
	"nbarecap/pkg/nba_api/models"
	"strings"
)

const (
	AbbrLen         = 3
	AbbrPairLen     = AbbrLen * 2
	SortKeySep      = "|"
	SplitSep        = "/"
	RsLinescore     = "LineScore"
	HGameId         = "GAME_ID"
	HStatusText     = "GAME_STATUS_TEXT"
	HGamecode       = "GAMECODE"
	HArenaName      = "ARENA_NAME"
	HNatlTv         = "NATL_TV_BROADCASTER_ABBREVIATION"
	HHomeTv         = "HOME_TV_BROADCASTER_ABBREVIATION"
	HAwayTv         = "AWAY_TV_BROADCASTER_ABBREVIATION"
	HHomeTeamId     = "HOME_TEAM_ID"
	HVisTeamId      = "VISITOR_TEAM_ID"
	HTeamId         = "TEAM_ID"
	HTeamAbbr       = "TEAM_ABBREVIATION"
	HTeamWinsLosses = "TEAM_WINS_LOSSES"
	HPts            = "PTS"
)

func fallbackAbbrFromGameCode(gameCode string) (away, home string) {
	tokens := strings.Split(gameCode, SplitSep)
	if len(tokens) != 2 {
		return "", ""
	}

	teams := tokens[1]
	if len(teams) == AbbrPairLen {
		return teams[:AbbrLen], teams[AbbrLen:]
	}
	return "", ""
}

func buildSortKey(status, gameID string) string {
	return status + SortKeySep + gameID
}

func BuildGameMap(responseMap map[string]any, ghRows []map[string]any) map[string]*models.GameInfo {
	gameMap := make(map[string]*models.GameInfo, len(ghRows))

	for _, r := range ghRows {
		id := utils.AnyAsString(r[HGameId])
		g := &models.GameInfo{
			GameID:   id,
			Status:   utils.AnyAsString(r[HStatusText]),
			GameCode: utils.AnyAsString(r[HGamecode]),
			Arena:    utils.AnyAsString(r[HArenaName]),
			NatTV:    utils.AnyAsString(r[HNatlTv]),
			HomeTV:   utils.AnyAsString(r[HHomeTv]),
			AwayTV:   utils.AnyAsString(r[HAwayTv]),
			HomeID:   utils.AnyAsString(r[HHomeTeamId]),
			AwayID:   utils.AnyAsString(r[HVisTeamId]),
		}
		g.AwayAbbr, g.HomeAbbr = fallbackAbbrFromGameCode(g.GameCode)
		g.SortKey = buildSortKey(g.Status, g.GameID)

		gameMap[id] = g
	}

	if lsAny, ok := responseMap[RsLinescore]; ok {
		if lsRows, ok := lsAny.([]map[string]any); ok {
			for _, r := range lsRows {
				gameID := utils.AnyAsString(r[HGameId])
				teamID := utils.AnyAsString(r[HTeamId])

				g := gameMap[gameID]
				if g == nil {
					continue
				}

				info := models.TeamInfo{
					Abbr:   utils.AnyAsString(r[HTeamAbbr]),
					Record: utils.AnyAsString(r[HTeamWinsLosses]),
					Pts:    utils.AnyAsIntPtr(r[HPts]),
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
