package go_sportradar_football

type Weather struct {
	TemperatureCelsius *int    `json:"temperature_celsius"`
	Wind               *string `json:"wind"`
	WindAdvantage      *string `json:"wind_advantage"`
	Pitch              *string `json:"pitch"`
	WeatherConditions  *string `json:"weather_conditions"`
}
