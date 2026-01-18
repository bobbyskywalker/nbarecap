package models

// TODO: name the fields readable

type BoxScoreTraditionalV3Response struct {
	BoxScoreTraditional BoxScoreTraditionalV3 `json:"boxScoreTraditional"`
}

type BoxScoreTraditionalV3 struct {
	GameID         string `json:"gameId"`
	GameCode       string `json:"gameCode,omitempty"`
	GameStatus     int    `json:"gameStatus,omitempty"`
	GameStatusText string `json:"gameStatusText,omitempty"`

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
	PersonID   int            `json:"personId"`
	FirstName  string         `json:"firstName,omitempty"`
	FamilyName string         `json:"familyName,omitempty"`
	NameI      string         `json:"nameI,omitempty"`
	Position   string         `json:"position,omitempty"`
	Comment    string         `json:"comment,omitempty"`
	JerseyNum  string         `json:"jerseyNum,omitempty"`
	Statistics map[string]any `json:"statistics,omitempty"`
}

type TeamV3 struct {
	TeamID      int            `json:"teamId"`
	TeamCity    string         `json:"teamCity,omitempty"`
	TeamName    string         `json:"teamName,omitempty"`
	TeamTricode string         `json:"teamTricode,omitempty"`
	Players     []PlayerV3     `json:"players,omitempty"`
	Statistics  map[string]any `json:"statistics,omitempty"`
	Starters    map[string]any `json:"starters,omitempty"`
	Bench       map[string]any `json:"bench,omitempty"`
}

type TeamStatsV3 struct {
	Minutes string `json:"minutes,omitempty"`

	Points int `json:"points,omitempty"`

	Fgm int     `json:"fgm,omitempty"`
	Fga int     `json:"fga,omitempty"`
	Fgp float64 `json:"fgp,omitempty"`

	Ftm int     `json:"ftm,omitempty"`
	Fta int     `json:"fta,omitempty"`
	Ftp float64 `json:"ftp,omitempty"`

	Tpm int     `json:"tpm,omitempty"`
	Tpa int     `json:"tpa,omitempty"`
	Tpp float64 `json:"tpp,omitempty"`

	OffReb int `json:"offReb,omitempty"`
	DefReb int `json:"defReb,omitempty"`
	TotReb int `json:"totReb,omitempty"`

	Ast int `json:"ast,omitempty"`
	Stl int `json:"stl,omitempty"`
	Blk int `json:"blk,omitempty"`

	Tov int `json:"tov,omitempty"`
	Pf  int `json:"pf,omitempty"`

	PlusMinus float64 `json:"plusMinus,omitempty"`
}

type PlayerStatsV3 struct {
	Points int `json:"points,omitempty"`

	Fgm int     `json:"fgm,omitempty"`
	Fga int     `json:"fga,omitempty"`
	Fgp float64 `json:"fgp,omitempty"`

	Ftm int     `json:"ftm,omitempty"`
	Fta int     `json:"fta,omitempty"`
	Ftp float64 `json:"ftp,omitempty"`

	Tpm int     `json:"tpm,omitempty"`
	Tpa int     `json:"tpa,omitempty"`
	Tpp float64 `json:"tpp,omitempty"`

	OffReb int `json:"offReb,omitempty"`
	DefReb int `json:"defReb,omitempty"`
	TotReb int `json:"totReb,omitempty"`

	Ast int `json:"ast,omitempty"`
	Stl int `json:"stl,omitempty"`
	Blk int `json:"blk,omitempty"`

	Tov int `json:"tov,omitempty"`
	Pf  int `json:"pf,omitempty"`

	PlusMinus float64 `json:"plusMinus,omitempty"`
}
