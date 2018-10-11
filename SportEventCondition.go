package go_sportradar_football

type SportEventCondition struct {
	Referee           Referee `json:"referee"`
	RefereeAssistants []struct {
		Type string `json:"type"`
		Referee
	} `json:"referee_assistants"`
	Venue       Venue   `json:"venue"`
	Attendance  int     `json:"attendance"`
	WeatherInfo Weather `json:"weather_info"`
}
