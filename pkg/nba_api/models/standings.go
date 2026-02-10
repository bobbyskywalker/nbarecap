package models

type Standing struct {
	LeagueID string `json:"LeagueID"`
	SeasonID string `json:"SeasonID"`
	TeamID   int64  `json:"TeamID"`

	TeamCity string `json:"TeamCity"`
	TeamName string `json:"TeamName"`
	TeamSlug string `json:"TeamSlug"`

	Conference       string `json:"Conference"`
	ConferenceRecord string `json:"ConferenceRecord"`
	PlayoffRank      int64  `json:"PlayoffRank"`
	ClinchIndicator  string `json:"ClinchIndicator"`
	Division         string `json:"Division"`
	DivisionRecord   string `json:"DivisionRecord"`
	DivisionRank     int64  `json:"DivisionRank"`

	Wins       int64   `json:"WINS"`
	Losses     int64   `json:"LOSSES"`
	WinPct     float64 `json:"WinPCT"`
	LeagueRank *int64  `json:"LeagueRank"`

	Record string `json:"Record"`
	Home   string `json:"HOME"`
	Road   string `json:"ROAD"`

	L10        string `json:"L10"`
	Last10Home string `json:"Last10Home"`
	Last10Road string `json:"Last10Road"`
	OT         string `json:"OT"`

	ThreePtsOrLess string `json:"ThreePTSOrLess"`
	TenPtsOrMore   string `json:"TenPTSOrMore"`

	LongHomeStreak    int64  `json:"LongHomeStreak"`
	StrLongHomeStreak string `json:"strLongHomeStreak"`
	LongRoadStreak    int64  `json:"LongRoadStreak"`
	StrLongRoadStreak string `json:"strLongRoadStreak"`
	LongWinStreak     int64  `json:"LongWinStreak"`
	LongLossStreak    int64  `json:"LongLossStreak"`

	CurrentHomeStreak    int64  `json:"CurrentHomeStreak"`
	StrCurrentHomeStreak string `json:"strCurrentHomeStreak"`
	CurrentRoadStreak    int64  `json:"CurrentRoadStreak"`
	StrCurrentRoadStreak string `json:"strCurrentRoadStreak"`
	CurrentStreak        int64  `json:"CurrentStreak"`
	StrCurrentStreak     string `json:"strCurrentStreak"`

	ConferenceGamesBack float64 `json:"ConferenceGamesBack"`
	DivisionGamesBack   float64 `json:"DivisionGamesBack"`

	ClinchedConferenceTitle int64 `json:"ClinchedConferenceTitle"`
	ClinchedDivisionTitle   int64 `json:"ClinchedDivisionTitle"`
	ClinchedPlayoffBirth    int64 `json:"ClinchedPlayoffBirth"`
	ClinchedPlayIn          int64 `json:"ClinchedPlayIn"`
	EliminatedConference    int64 `json:"EliminatedConference"`
	EliminatedDivision      int64 `json:"EliminatedDivision"`

	AheadAtHalf  string `json:"AheadAtHalf"`
	BehindAtHalf string `json:"BehindAtHalf"`
	TiedAtHalf   string `json:"TiedAtHalf"`

	AheadAtThird  string `json:"AheadAtThird"`
	BehindAtThird string `json:"BehindAtThird"`
	TiedAtThird   string `json:"TiedAtThird"`

	Score100PTS    string `json:"Score100PTS"`
	OppScore100PTS string `json:"OppScore100PTS"`
	OppOver500     string `json:"OppOver500"`
	LeadInFGPCT    string `json:"LeadInFGPCT"`
	LeadInReb      string `json:"LeadInReb"`
	FewerTurnovers string `json:"FewerTurnovers"`

	PointsPG     float64 `json:"PointsPG"`
	OppPointsPG  float64 `json:"OppPointsPG"`
	DiffPointsPG float64 `json:"DiffPointsPG"`

	VsEast      string `json:"vsEast"`
	VsAtlantic  string `json:"vsAtlantic"`
	VsCentral   string `json:"vsCentral"`
	VsSoutheast string `json:"vsSoutheast"`
	VsWest      string `json:"vsWest"`
	VsNorthwest string `json:"vsNorthwest"`
	VsPacific   string `json:"vsPacific"`
	VsSouthwest string `json:"vsSouthwest"`

	Jan string `json:"Jan"`
	Feb string `json:"Feb"`
	Mar string `json:"Mar"`
	Apr string `json:"Apr"`
	May string `json:"May"`
	Jun string `json:"Jun"`
	Jul string `json:"Jul"`
	Aug string `json:"Aug"`
	Sep string `json:"Sep"`
	Oct string `json:"Oct"`
	Nov string `json:"Nov"`
	Dec string `json:"Dec"`

	Score80Plus     string `json:"Score_80_Plus"`
	OppScore80Plus  string `json:"Opp_Score_80_Plus"`
	ScoreBelow80    string `json:"Score_Below_80"`
	OppScoreBelow80 string `json:"Opp_Score_Below_80"`

	TotalPoints     int64 `json:"TotalPoints"`
	OppTotalPoints  int64 `json:"OppTotalPoints"`
	DiffTotalPoints int64 `json:"DiffTotalPoints"`

	LeagueGamesBack    float64 `json:"LeagueGamesBack"`
	PlayoffSeeding     int64   `json:"PlayoffSeeding"`
	ClinchedPostSeason int64   `json:"ClinchedPostSeason"`

	Neutral string `json:"NEUTRAL"`
}
