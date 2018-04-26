package go_sportradar_football

type DailyResult struct {
	Meta
	Results []struct {
		SportEvent       SportEvent       `json:"sport_event"`
		SportEventStatus SportEventStatus `json:"sport_event_status"`
	} `json:"results"`
}

type LiveResult struct {
	Meta
	Results []struct {
		SportEvent       SportEvent       `json:"sport_event"`
		SportEventStatus SportEventStatus `json:"sport_event_status"`
	} `json:"results"`
}

type DailySchedule struct {
	Meta
	SportEvent []SportEvent `json:"sport_events"`
}

type Schedule struct {
	Meta
	Tournament Tournament `json:"tournament"`
	Category   Category   `json:"category"`

	SportEvent []SportEvent `json:"sport_events"`
}

type TournamentRound struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
}

type Competitor struct {
	BaseTeam
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	Abbreviation string `json:"abbreviation"`
	Qualifier    string `json:"qualifier"`
}

type Venue struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Capacity       int    `json:"capacity"`
	CityName       string `json:"city_name"`
	CountryName    string `json:"country_name"`
	MapCoordinates string `json:"map_coordinates"`
	CountryCode    string `json:"country_code"`
}
