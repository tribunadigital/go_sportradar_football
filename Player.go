package go_sportradar_football

type BasePlayer struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	JerseyNumber  int    `json:"jersey_number"`
}

type Player struct {
	BasePlayer
	DateOfBirth   string `json:"date_of_birth"`
	Nationality   string `json:"nationality"`
	CountryCode   string `json:"country_code"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	PreferredFoot string `json:"preferred_foot"`
	Gender        string `json:"gender"`
}
