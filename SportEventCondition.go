package go_sportradar_football

type SportEventCondition struct {
	Referee           Referee `json:"referee"`
	RefereeAssistants []struct {
		Type string `json:"type"`
		Referee
	} `json:"referee_assistants"`
	Venue       Venue `json:"venue"`
	Attendance  int   `json:"attendance"`
	WeatherInfo struct {
		Pitch             string `json:"pitch"`
		WeatherConditions string `json:"weather_conditions"`
	} `json:"weather_info"`
}
