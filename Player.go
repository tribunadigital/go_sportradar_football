package go_sportradar_football

type BasePlayer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	JerseyNumber int    `json:"jersey_number"`
}

type Player struct {
	BasePlayer
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Nationality   string `json:"nationality"`
	CountryCode   string `json:"country_code"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	PreferredFoot string `json:"preferred_foot"`
	Gender        string `json:"gender"`
}

type Role struct {
	Type          string     `json:"type"`
	Active        string     `json:"active"`
	Team          DetailTeam `json:"team"`
	StartDate     string     `json:"start_date"`
	EndDate       string     `json:"end_date"`
	JerseyNumber  string     `json:"jersey_number"`
	PreferredFoot string     `json:"preferred_foot"`
	Gender        string     `json:"gender"`
}

type PlayerProfile struct {
	Meta
	Player Player       `json:"player"`
	Teams  []DetailTeam `json:"teams"`
	Roles  []Role       `json:"roles"`
}
