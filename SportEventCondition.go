package go_sportradar_football

type SportEventCondition struct {
	Referee struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		BaseNational
	} `json:"referee"`
	RefereeAssistants []struct {
		Type string `json:"type"`
		Id   string `json:"id"`
		Name string `json:"name"`
		BaseNational
	}
	Venue      Venue `json:"venue"`
	Attendance int   `json:"attendance"`
	WeatherInfo struct {
		Pitch             string `json:"pitch"`
		WeatherConditions string `json:"weather_conditions"`
	} `json:"weather_info"`
}
