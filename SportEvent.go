package go_sportradar_football

type MatchLineups struct {
	Meta
	SportEvent SportEvent `json:"sport_event"`
	Lineups    []Lineups  `json:"lineups"`
}

type MatchSummary struct {
	Meta
	SportEvent SportEvent `json:"sport_event"`
	SportEventConditions struct {
		//Referee
		Venue      Venue `json:"venue"`
		Attendance int   `json:"attendance"`
	} `json:"sport_event_conditions"`
	SportEventStatus SportEventStatus `json:"sport_event_status"`
	Statistics struct {
		Teams []Team `json:"teams"`
	} `json:"statistics"`
}

type SportEvent struct {
	Id              string          `json:"id"`
	Scheduled       string          `json:"scheduled"`
	StartTimeTbd    bool            `json:"start_time_tbd"`
	Status          string          `json:"status"`
	TournamentRound TournamentRound `json:"tournament_round"`
	Season          Season          `json:"season"`
	Tournament      Tournament      `json:"tournament"`

	Competitors []Competitor `json:"competitors"`
	Venue       Venue        `json:"venue"`
}

type PeriodScore struct {
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
	Type      string `json:"type"`
	Number    int    `json:"number"`
}

type SportEventStatus struct {
	Status      string `json:"status"`
	MatchStatus string `json:"match_status"`
	HomeScore   int    `json:"home_score"`
	AwayScore   int    `json:"away_score"`
	WinnerId    string `json:"winner_id"`

	PeriodScores []PeriodScore `json:"period_scores"`
}

type Lineups struct {
	Team      string `json:"team"`
	Formation string `json:"formation"`
	Manager struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
		CountryCode string `json:"country_code"`
	} `json:"manager"`
	Jersey struct {
		Type              string `json:"type"`
		Base              string `json:"base"`
		Sleeve            string `json:"sleeve"`
		Number            string `json:"number"`
		Squares           bool   `json:"squares"`
		Stripes           bool   `json:"stripes"`
		StripesColor      string `json:"stripes_color"`
		HorizontalStripes bool   `json:"horizontal_stripes"`
		Split             bool   `json:"split"`
		ShirtType         string `json:"shirt_type"`
		SleeveDetail      string `json:"sleeve_detail"`
	} `json:"jersey"`

	StartingLineup []StartingLineup `json:"starting_lineup"`

	Substitutes []struct {
		BasePlayer
	} `json:"substitutes"`
}

type StartingLineup struct {
	BasePlayer
	Position string `json:"position"`
	Order    int    `json:"order"`
}
