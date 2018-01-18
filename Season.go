package go_sportradar_football



type Season struct {
	Id string `json:"id"`
	Name string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Year string `json:"year"`
}