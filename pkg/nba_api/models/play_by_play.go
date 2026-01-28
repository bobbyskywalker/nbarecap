package models

import "time"

type PlayByPlayResponse struct {
	Meta Meta `json:"meta"`
	Game Game `json:"game"`
}

type Meta struct {
	Version int       `json:"version"`
	Request string    `json:"request"`
	Time    time.Time `json:"time"`
}

type Game struct {
	GameID         string   `json:"gameId"`
	VideoAvailable int      `json:"videoAvailable"`
	Actions        []Action `json:"actions"`
}

type Action struct {
	ActionNumber int    `json:"actionNumber"`
	Clock        string `json:"clock"`
	Period       int    `json:"period"`

	TeamID      int    `json:"teamId"`
	TeamTricode string `json:"teamTricode"`

	PersonID    int    `json:"personId"`
	PlayerName  string `json:"playerName"`
	PlayerNameI string `json:"playerNameI"`

	XLegacy int `json:"xLegacy"`
	YLegacy int `json:"yLegacy"`

	ShotDistance int    `json:"shotDistance"`
	ShotResult   string `json:"shotResult"`
	IsFieldGoal  int    `json:"isFieldGoal"`

	ScoreHome   string `json:"scoreHome"`
	ScoreAway   string `json:"scoreAway"`
	PointsTotal int    `json:"pointsTotal"`

	Location    string `json:"location"`
	Description string `json:"description"`

	ActionType string `json:"actionType"`
	SubType    string `json:"subType"`

	VideoAvailable int `json:"videoAvailable"`
	ShotValue      int `json:"shotValue"`
	ActionID       int `json:"actionId"`
}
