package nba

const (
	Title = " NBA games\n\n"

	DateFormat = "2006-01-02"

	AbbrLen     = 3
	AbbrPairLen = AbbrLen * 2

	SortKeySep = "|"
	SplitSep   = "/"

	ScoresFormat = " — %d-%d"

	GameInfoFormat = "%s — %s%s @ %s%s%s\n"

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

	NationalBroadcastInfoFormat = "National: %s"
	AwayBroadcastInfoFormat     = "Away: %s"
	HomeBroadcastInfoFormat     = "Home: %s"

	HTeamId         = "TEAM_ID"
	HTeamAbbr       = "TEAM_ABBREVIATION"
	HTeamWinsLosses = "TEAM_WINS_LOSSES"
	HPts            = "PTS"
)
