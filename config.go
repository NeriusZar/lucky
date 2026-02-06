package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/NeriusZar/lucky/internal/database"
	"github.com/NeriusZar/lucky/internal/luck"
	"github.com/NeriusZar/lucky/internal/meteoapi"
	_ "github.com/lib/pq"
)

type config struct {
	tick time.Duration
	api  *meteoapi.ApiClient
	db   *database.Queries
	lc   luck.LuckCalculator
}

func (c *config) init() error {
	defaultTickParam, ok := os.LookupEnv("DEFAULT_TICK_SECONDS")
	if !ok {
		return fmt.Errorf("DEFAULT_TICK_SECONDS was not provided in .env file")
	}

	defaultTickInSeconds, err := strconv.Atoi(defaultTickParam)
	if err != nil {
		return fmt.Errorf("Failed to parse DEFAULT_TICK_SECONDS to integer. %v", err)
	}

	dbUrl, ok := os.LookupEnv("POSTGRESQL_URL")
	if !ok {
		return fmt.Errorf("POSTGRESQL_URL was not provided in .env file")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to open database. %v", err)
	}
	dbQueries := database.New(db)

	api := meteoapi.NewApiClient()
	lc := luck.NewLuckCalculator(dbQueries)
	defaultTick := time.Duration(defaultTickInSeconds) * time.Second
	c.tick = defaultTick
	c.api = &api
	c.db = dbQueries
	c.lc = lc

	return nil
}
