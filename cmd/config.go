package cmd

import (
	"github.com/NeriusZar/lucky/internal/database"
	"github.com/NeriusZar/lucky/internal/meteoapi"
)

type Config struct {
	api meteoapi.ApiClient
	db  *database.Queries
}

func NewConfig(db *database.Queries) Config {
	return Config{
		api: meteoapi.NewApiClient(),
		db:  db,
	}
}
