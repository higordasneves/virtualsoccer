package soccer

type Fixture struct {
	DateDesc      string      `json:"dateDesc"`
	Desc          string      `json:"desc"`
	FixtureID     int         `json:"fixtureId"`
	ChallengeID   int         `json:"challengeId"`
	CompetitionID int         `json:"competitionId"`
	Title         interface{} `json:"title"`
	Time          string      `json:"time"`
	UserTime      string      `json:"userTime"`
}

type Matches struct {
	Title    string    `json:"title"`
	Fixtures []Fixture `json:"fixtures"`
}

type MatchResult struct {
	Date         string   `json:"date"`
	Title        string   `json:"title"`
	MatchMarkets []string `json:"matchMarkets"`
}

type MatchResultReport struct {
	Date                    string
	Title                   string
	FinalResult             string
	ExactScore              string
	ExactGoalNumber         string
	OverHalfAGoal           bool
	UnderHalfAGoal          bool
	OverOneAndAHalfAGoal    bool
	UnderOneAndAHalfAGoal   bool
	OverTwoAndAHalfAGoal    bool
	UnderTwoAndAHalfAGoal   bool
	OverThreeAndAHalfAGoal  bool
	UnderThreeAndAHalfAGoal bool
	HomeTeamGoals           string
	InvitedTeamGoals        string
	BothTeamsToScore        bool
	HomeTeamToScore         bool
	InvitedTeamToScore      bool
}
