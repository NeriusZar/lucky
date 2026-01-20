package cmd

import (
	"github.com/NeriusZar/lucky/internal/meteoapi"
	"github.com/NeriusZar/lucky/internal/persistance"
)

type Config struct {
	api meteoapi.ApiClient
	db  persistance.LuckyDB
}

func NewConfig() Config {
	return Config{
		api: meteoapi.NewApiClient(),
		db:  persistance.NewLuckyDB(),
	}
}
