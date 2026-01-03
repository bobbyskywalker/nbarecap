package recaps

const (
	TITLE = "Today's NBA games\n\n"

	AbbrLen     = 3
	AbbrPairLen = AbbrLen * 2

	SortKeySep = "|"
	SplitSep   = "/"

	ScoresFormat = " — %d-%d"

	HeaderMissingGameheader = "No GameHeader in response.\n"
	HeaderBadGameheader     = "GameHeader has unexpected type.\n"

	RsGameheader = "GameHeader"
	RsLinescore  = "LineScore"

	HGameId     = "GAME_ID"
	HStatusText = "GAME_STATUS_TEXT"
	HGamecode   = "GAMECODE"
	HArenaName  = "ARENA_NAME"
	HNatlTv     = "NATL_TV_BROADCASTER_ABBREVIATION"
	HHomeTv     = "HOME_TV_BROADCASTER_ABBREVIATION"
	HAwayTv     = "AWAY_TV_BROADCASTER_ABBREVIATION"
	HHomeTeamId = "HOME_TEAM_ID"
	HVisTeamId  = "VISITOR_TEAM_ID"

	HTeamId         = "TEAM_ID"
	HTeamAbbr       = "TEAM_ABBREVIATION"
	HTeamWinsLosses = "TEAM_WINS_LOSSES"
	HPts            = "PTS"
)
