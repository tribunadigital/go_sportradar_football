package go_sportradar_football

type BaseTeam struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DetailTeam struct {
	BaseTeam
	Country      string   `json:"country"`
	CountryCode  string   `json:"country_code"`
	Category     Category `json:"category"`
	Abbreviation string   `json:"abbreviation"`
}

type TeamStanding struct {
	Team           BaseTeam `json:"team"`
	Rank           int      `json:"rank"`
	CurrentOutcome string   `json:"current_outcome"`
	GroupName      string   `json:"group_name"`
	Played         int      `json:"played"`
	Win            int      `json:"win"`
	Draw           int      `json:"draw"`
	Loss           int      `json:"loss"`
	GoalsFor       int      `json:"goals_for"`
	GoalsAgainst   int      `json:"goals_against"`
	GoalDiff       int      `json:"goal_diff"`
	Points         int      `json:"points"`
}

type TeamProfile struct {
	Team DetailTeam `json:"team"`

	Venue   Venue    `json:"venue"`
	Players []Player `json:"players"`
	Statistics struct {
		Seasons []SeasonStat `json:"seasons"`
	} `json:"statistics"`
}
