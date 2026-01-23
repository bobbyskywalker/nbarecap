package models

type BoxScoreTraditionalV3Response struct {
	BoxScoreTraditional BoxScoreTraditionalV3 `json:"boxScoreTraditional"`
}

type BoxScoreTraditionalV3 struct {
	GameID         string `json:"gameId"`
	GameCode       string `json:"gameCode,omitempty"`
	GameStatus     int    `json:"gameStatus,omitempty"`
	GameStatusText string `json:"gameStatusText,omitempty"`

	AwayTeamID int `json:"awayTeamId,omitempty"`
	HomeTeamID int `json:"homeTeamId,omitempty"`

	Period   PeriodV3 `json:"period,omitempty"`
	HomeTeam TeamV3   `json:"homeTeam"`
	AwayTeam TeamV3   `json:"awayTeam"`
}

type PeriodV3 struct {
	Current       int  `json:"current,omitempty"`
	Type          int  `json:"type,omitempty"`
	MaxRegular    int  `json:"maxRegular,omitempty"`
	IsHalftime    bool `json:"isHalftime,omitempty"`
	IsEndOfPeriod bool `json:"isEndOfPeriod,omitempty"`
}

type GameLeadersV3 struct {
	HomeLeaders LeadersSideV3 `json:"homeLeaders,omitempty"`
	AwayLeaders LeadersSideV3 `json:"awayLeaders,omitempty"`
}

type LeadersSideV3 struct {
	PersonID  int    `json:"personId,omitempty"`
	Name      string `json:"name,omitempty"`
	JerseyNum string `json:"jerseyNum,omitempty"`
	TeamID    int    `json:"teamId,omitempty"`
	Points    int    `json:"points,omitempty"`
	Rebounds  int    `json:"rebounds,omitempty"`
	Assists   int    `json:"assists,omitempty"`
}

type PlayerV3 struct {
	PersonID   int    `json:"personId"`
	FirstName  string `json:"firstName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	NameI      string `json:"nameI,omitempty"`

	PlayerSlug string `json:"playerSlug,omitempty"`

	Position  string `json:"position,omitempty"`
	Comment   string `json:"comment,omitempty"`
	JerseyNum string `json:"jerseyNum,omitempty"`

	Statistics PlayerStatsV3 `json:"statistics,omitempty"`
}

type TeamV3 struct {
	TeamID      int    `json:"teamId"`
	TeamCity    string `json:"teamCity,omitempty"`
	TeamName    string `json:"teamName,omitempty"`
	TeamTricode string `json:"teamTricode,omitempty"`
	TeamSlug    string `json:"teamSlug,omitempty"`

	Players []PlayerV3 `json:"players,omitempty"`

	Statistics TeamStatsV3 `json:"statistics,omitempty"`
	Starters   TeamStatsV3 `json:"starters,omitempty"`
	Bench      TeamStatsV3 `json:"bench,omitempty"`
}

type TeamStatsV3 = PlayerStatsV3

type PlayerStatsV3 struct {
	Minutes string `json:"minutes,omitempty"`

	FieldGoalsMade       int     `json:"fieldGoalsMade,omitempty"`
	FieldGoalsAttempted  int     `json:"fieldGoalsAttempted,omitempty"`
	FieldGoalsPercentage float64 `json:"fieldGoalsPercentage,omitempty"`

	ThreePointersMade       int     `json:"threePointersMade,omitempty"`
	ThreePointersAttempted  int     `json:"threePointersAttempted,omitempty"`
	ThreePointersPercentage float64 `json:"threePointersPercentage,omitempty"`

	FreeThrowsMade       int     `json:"freeThrowsMade,omitempty"`
	FreeThrowsAttempted  int     `json:"freeThrowsAttempted,omitempty"`
	FreeThrowsPercentage float64 `json:"freeThrowsPercentage,omitempty"`

	ReboundsOffensive int `json:"reboundsOffensive,omitempty"`
	ReboundsDefensive int `json:"reboundsDefensive,omitempty"`
	ReboundsTotal     int `json:"reboundsTotal,omitempty"`

	Assists       int `json:"assists,omitempty"`
	Steals        int `json:"steals,omitempty"`
	Blocks        int `json:"blocks,omitempty"`
	Turnovers     int `json:"turnovers,omitempty"`
	FoulsPersonal int `json:"foulsPersonal,omitempty"`

	Points          int     `json:"points,omitempty"`
	PlusMinusPoints float64 `json:"plusMinusPoints,omitempty"`
}
