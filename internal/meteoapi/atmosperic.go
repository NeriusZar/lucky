package meteoapi

import (
	"strings"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (api ApiClient) GetCurrentAtmosphericData(lat, long float64) (CurrentWeatherResponse, error) {
	u, err := url.Parse(baseUrl + forecasePath)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	q := u.Query()
	q.Add("latitude", fmt.Sprintf("%f", lat))
	q.Add("longitude", fmt.Sprintf("%f", long))
	q.Add("current", getWeatherVariables())
	q.Add("wind_speed_unit", "ms")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	res, err := api.client.Do(req)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return CurrentWeatherResponse{}, fmt.Errorf("Failed to fetch current weather data. Status code - %d", res.StatusCode)
	}
	defer res.Body.Close()

	var currentWeatherResponse CurrentWeatherResponse
	if err := json.NewDecoder(res.Body).Decode(&currentWeatherResponse); err != nil {
		return CurrentWeatherResponse{}, err
	}

	return currentWeatherResponse, nil
}

func getWeatherVariables() string {
	variables := []string{
		"temperature_2m",
		"wind_speed_10m",
		"cloud_cover",
		"pressure_msl",
	}

	var result strings.Builder
	for i, v := range variables {
		if i == len(variables)-1 {
			result.WriteString(v)
			break
		}
		result.WriteString(fmt.Sprintf("%s,", v))
	}

	return result.String()
}
