package meteoapi

type CurrentWeatherResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentUnits         struct {
		Time          string `json:"time"`
		Interval      string `json:"interval"`
		Temperature2M string `json:"temperature_2m"`
		WindSpeed10M  string `json:"wind_speed_10m"`
		CloudCover    string `json:"cloud_cover"`
		PressureMsl   string `json:"pressure_msl"`
	} `json:"current_units"`
	Current struct {
		Time          string  `json:"time"`
		Interval      int     `json:"interval"`
		Temperature2M float64 `json:"temperature_2m"`
		WindSpeed10M  float64 `json:"wind_speed_10m"`
		CloudCover    int     `json:"cloud_cover"`
		PressureMsl   float64 `json:"pressure_msl"`
	} `json:"current"`
}
