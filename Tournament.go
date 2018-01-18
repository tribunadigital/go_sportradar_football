package go_sportradar_football




type Tournament struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Sport Sport `json:"sport"`
	Category Category `json:"category"`
	CurrentSeason Season `json:"current_season"`
}


type TournamentsResult struct {
	Meta
	Tournaments []Tournament `json:"tournaments"`
}