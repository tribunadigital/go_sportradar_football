package go_sportradar_football

type BasePlayer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	JerseyNumber int    `json:"jersey_number"`
}

type Player struct {
	BasePlayer
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Nationality   string `json:"nationality"`
	CountryCode   string `json:"country_code"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	PreferredFoot string `json:"preferred_foot"`
	Gender        string `json:"gender"`
}

type Role struct {
	Type          string     `json:"type"`
	Active        string     `json:"active"`
	Team          DetailTeam `json:"team"`
	StartDate     string     `json:"start_date"`
	EndDate       string     `json:"end_date"`
	JerseyNumber  int        `json:"jersey_number"`
	PreferredFoot string     `json:"preferred_foot"`
	Gender        string     `json:"gender"`
}

type PlayerProfile struct {
	Meta
	Player Player       `json:"player"`
	Teams  []DetailTeam `json:"teams"`
	Roles  []Role       `json:"roles"`
	Statistics struct {
		Seasons []struct {
			Id         string           `json:"id"`
			Name       string           `json:"name"`
			Team       DetailTeam       `json:"team"`
			Tournament Tournament       `json:"tournament"`
			Statistics PlayerSeasonStat `json:"statistics"`
		} `json:"seasons"`
		Total PlayerSeasonStat `json:"total"`
	} `json:"statistics"`
}

type PlayerSeasonStat struct {
	MatchesPlayed  int    `json:"matches_played"`
	SubstitutedIn  int    `json:"substituted_in"`
	GoalScored     int    `json:"goal_scored"`
	Assists        int    `json:"assists"`
	OwnGoals       int    `json:"own_goals"`
	YellowCards    int    `json:"yellow_cards"`
	YellowRedCards int    `json:"yellow_red_cards"`
	RedCards       int    `json:"red_cards"`
	LastEventTime  string `json:"last_event_time"`
}
