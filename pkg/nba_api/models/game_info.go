package models

type TeamInfo struct {
	Abbr   string
	Record string
	Pts    *int
}

type GameInfo struct {
	GameID   string
	Status   string
	GameCode string
	Arena    string
	NatTV    string
	HomeTV   string
	AwayTV   string
	HomeID   string
	AwayID   string
	Home     TeamInfo
	Away     TeamInfo
	HomeAbbr string
	AwayAbbr string
	SortKey  string
}

type GameInfoFormatted struct {
	GameId   string
	GameInfo string
}

func NewGameInfoFormatted(gameId string, gameInfo string) GameInfoFormatted {
	return GameInfoFormatted{GameId: gameId, GameInfo: gameInfo}
}
