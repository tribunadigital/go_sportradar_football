package go_sportradar_football

type MatchStatistics struct {
	FreeKicks       string `json:"free_kicks"`
	Offsides        *int   `json:"offsides"`
	ShotsOffTarget  *int   `json:"shots_off_target"`
	CornerKicks     *int   `json:"corner_kicks"`
	Injuries        *int   `json:"injuries"`
	ShotsOnTarget   *int   `json:"shots_on_target"`
	YellowCards     *int   `json:"yellow_cards"`
	YellowRedCards  *int   `json:"yellow_red_cards"`
	RedCards        *int   `json:"red_cards"`
	BallPossession  *int   `json:"ball_possession"`
	GoalKicks       *int   `json:"goal_kicks"`
	ThrowIns        *int   `json:"throw_ins"`
	ShotsSaved      *int   `json:"shots_saved"`
	Fouls           *int   `json:"fouls"`
	OwnGoals        *int   `json:"own_goals"`
	PenaltiesMissed *int   `json:"penalties_missed"`
	ShotsBlocked    *int   `json:"shots_blocked"`
}

type PlayerStatistics struct {
	SubstitutedIn          *int `json:"substituted_in"`
	SubstitutedOut         *int `json:"substituted_out"`
	GoalsScored            *int `json:"goals_scored"`
	Assists                *int `json:"assists"`
	OwnGoals               *int `json:"own_goals"`
	YellowCards            *int `json:"yellow_cards"`
	YellowRedCards         *int `json:"yellow_red_cards"`
	RedCards               *int `json:"red_cards"`
	PenaltiesMissed        *int `json:"penalties_missed"`
	PenaltyGoalsScored     *int `json:"penalty_goals_scored"`
	GoalLineClearances     *int `json:"goal_line_clearances"`
	Interceptions          *int `json:"interceptions"`
	ChancesCreated         *int `json:"chances_created"`
	CrossesSuccessful      *int `json:"crosses_successful"`
	CrossesTotal           *int `json:"crosses_total"`
	PassesShortSuccessful  *int `json:"passes_short_successful"`
	PassesMediumSuccessful *int `json:"passes_medium_successful"`
	PassesLongSuccessful   *int `json:"passes_long_successful"`
	PassesShortTotal       *int `json:"passes_short_total"`
	PassesMediumTotal      *int `json:"passes_medium_total"`
	PassesLongTotal        *int `json:"passes_long_total"`
	DuelsHeaderSuccessful  *int `json:"duels_header_successful"`
	DuelsSprintSuccessful  *int `json:"duels_sprint_successful"`
	DuelsTackleSuccessful  *int `json:"duels_tackle_successful"`
	DuelsSprintTotal       *int `json:"duels_sprint_total"`
	DuelsTackleTotal       *int `json:"duels_tackle_total"`
	DuelsHeaderTotal       *int `json:"duels_header_total"`
	GoalsConceded          *int `json:"goals_conceded"`
	ShotsFacedSaved        *int `json:"shots_faced_saved"`
	ShotsFacedTotal        *int `json:"shots_faced_total"`
	FoulsCommitted         *int `json:"fouls_committed"`
	WasFouled              *int `json:"was_fouled"`
	Offsides               *int `json:"offsides"`
	ShotsOnGoal            *int `json:"shots_on_goal"`
	ShotsOffGoal           *int `json:"shots_off_goal"`
	ShotsBlocked           *int `json:"shots_blocked"`
	MinutesPlayed          *int `json:"minutes_played"`
	PerformanceScore       *int `json:"performance_score"`
	GoalsByHead            *int `json:"goals_by_head"`
	GoalAttempts           *int `json:"goal_attempts"`
	PenaltiesFaced         *int `json:"penalties_faced"`
	PenaltiesSaved         *int `json:"penalties_saved"`
	GoalsByPenalty         *int `json:"goals_by_penalty"`
}
