package go_sportradar_football

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TeamStanding struct {
	Team Team `json:"team"`

	Rank           int    `json:"rank"`
	CurrentOutcome string `json:"current_outcome"`
	Played         int    `json:"played"`
	Win            int    `json:"win"`
	Draw           int    `json:"draw"`
	Loss           int    `json:"loss"`
	GoalsFor       int    `json:"goals_for"`
	GoalsAgainst   int    `json:"goals_against"`
	GoalDiff       int    `json:"goal_diff"`
	Points         int    `json:"points"`
}

type TeamProfile struct {
	Team     Team     `json:"team"`
	Category Category `json:"category"`
	Venue    Venue    `json:"venue"`
	Players  []Player `json:"players"`
}
