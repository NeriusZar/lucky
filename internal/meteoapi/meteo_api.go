package meteoapi

import (
	"net/http"
)

type ApiClient struct {
	client http.Client
}

const baseUrl = "https://api.open-meteo.com"
const forecasePath = "/v1/forecast"

func NewApiClient() ApiClient {
	return ApiClient{
		http.Client{},
	}
}
