package go_sportradar_football

type MatchLineups struct {
	Meta
	SportEvent SportEvent `json:"sport_event"`
	Lineups    []Lineups  `json:"lineups"`
}

type MatchTimeline struct {
	Meta
	SportEvent          SportEvent          `json:"sport_event"`
	SportEventCondition SportEventCondition `json:"sport_event_conditions"`
	SportEventStatus    SportEventStatus    `json:"sport_event_status"`
	Timeline            []TimeLineEvent     `json:"timeline"`
	Statistics          struct {
		Teams []struct {
			BaseTeam
			Abbreviation string          `json:"abbreviation"`
			Qualifier    string          `json:"qualifier"`
			Statistics   MatchStatistics `json:"statistics"`
			Players      []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				PlayerStatistics
			} `json:"players"`
		} `json:"teams"`
	} `json:"statistics"`
}

type TimelinePlayer struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Type *string `json:"type"`
}

type TimeLineEvent struct {
	Id                int             `json:"id"`
	Type              string          `json:"type"`
	Time              string          `json:"time"`
	PeriodName        *string         `json:"period_name"`
	MatchTime         *int            `json:"match_time"`
	MatchClock        *string         `json:"match_clock"`
	Team              *string         `json:"team"`
	X                 *int            `json:"x"`
	Y                 *int            `json:"y"`
	HomeScore         *int            `json:"home_score"`
	AwayScore         *int            `json:"away_score"`
	ShootoutHomeScore *int            `json:"shootout_home_score"`
	ShootoutAwayScore *int            `json:"shootout_away_score"`
	Period            *int            `json:"period"`
	PeriodType        *string         `json:"period_type"`
	Status            *string         `json:"status"`
	StoppageTime      *string         `json:"stoppage_time"`
	Outcome           *string         `json:"outcome"`
	GoalScorer        *TimelinePlayer `json:"goal_scorer"`
	Assist            *TimelinePlayer `json:"assist"`
	Player            *TimelinePlayer `json:"player"`
	PlayerOut         *TimelinePlayer `json:"player_out"`
	PlayerIn          *TimelinePlayer `json:"player_in"`
}

type MatchSummary struct {
	Meta
	SportEvent           SportEvent `json:"sport_event"`
	SportEventConditions struct {
		//Referee
		Venue      Venue `json:"venue"`
		Attendance int   `json:"attendance"`
	} `json:"sport_event_conditions"`
	SportEventStatus SportEventStatus `json:"sport_event_status"`
	Statistics       struct {
		Teams []BaseTeam `json:"teams"`
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
	Manager   struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		BaseNational
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
