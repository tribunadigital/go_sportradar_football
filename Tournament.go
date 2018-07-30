package go_sportradar_football

type Tournament struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Sport         Sport    `json:"sport"`
	Category      Category `json:"category"`
	CurrentSeason Season   `json:"current_season"`
}

type TournamentsResult struct {
	Meta
	Tournaments []Tournament `json:"tournaments"`
}

type TournamentSeason struct {
	Meta
	Tournament       `json:"tournament"`
	Seasons []Season `json:"seasons"`
}

type Group struct {
	Name         string         `json:"name"`
	TeamStanding []TeamStanding `json:"team_standings"`
}

type Standing struct {
	TieBreakRule string  `json:"tie_break_rule"`
	Type         string  `json:"type"`
	Group        []Group `json:"groups"`
}

type TournamentStanding struct {
	Meta
	Tournament          `json:"tournament"`
	Season   Season     `json:"season"`
	Standing []Standing `json:"standings"`
}
