package go_sportradar_football

type Season struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Year      string `json:"year"`
}

type SeasonStat struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Statistics struct {
		MatchesPlayed int    `json:"matches_played"`
		MatchesWon    int    `json:"matches_won"`
		MatchesDrawn  int    `json:"matches_drawn"`
		MatchesLost   int    `json:"matches_lost"`
		GoalsScored   int    `json:"goals_scored"`
		GoalsConceded int    `json:"goals_conceded"`
		GroupPosition int    `json:"group_position"`
		CupRank       int    `json:"cup_rank"`
		GroupName     string `json:"group_name"`
	} `json:"statistics"`
	Tournament Tournament `json:"tournament"`
	Form struct {
		Total string `json:"total"`
		Home  string `json:"home"`
		Away  string `json:"away"`
	} `json:"form"`
}
