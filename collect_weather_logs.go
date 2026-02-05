package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/NeriusZar/lucky/internal/database"
)

func collectWeahterLogs(ctx context.Context, c *config, cmd command) error {
	collectCmd := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	outputFilePath := collectCmd.String("out", "", "redirects output to the provided file")
	tick := collectCmd.Duration("tick", time.Second*60, "overrides default tick duration for weather data collection")
	help := collectCmd.Bool("help", false, "gives instructions about the collect command")

	if err := collectCmd.Parse(cmd.Args); err != nil {
		return fmt.Errorf("failed to parse flags %v", err)
	}

	if help != nil && *help {
		fmt.Printf("usage: collect [--out] [--tick]\n")
		return nil
	}

	log.SetOutput(os.Stdout)

	if outputFilePath != nil && *outputFilePath != "" {
		filePath := filepath.Join(*outputFilePath)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create/open a file. %v", err)
		}
		defer file.Close()

		log.SetOutput(file)
	}

	tickDuration := c.tick
	if tick != nil {
		tickDuration = *tick
	}

	log.Print("starting to collect weather data")

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(tickDuration):
			if err := processWeatherData(ctx, c); err != nil {
				fmt.Printf("failed to collect weather logs. %v", err)
			}
		}
	}
}

func processWeatherData(ctx context.Context, c *config) error {
	log.Printf("checking weather on %v\n", time.Now().UTC())

	locations, err := c.db.GetAllLocations(ctx)
	if err != nil {
		return fmt.Errorf("failed to get locations %v", err)
	}

	for _, l := range locations {
		if err := fetchAndStoreWeatherLog(ctx, c, l); err != nil {
			return err
		}
	}

	return nil
}

func fetchAndStoreWeatherLog(ctx context.Context, c *config, location database.Location) error {
	weatherRes, err := c.api.GetCurrentAtmosphericData(ctx, location.Latitude, location.Longitude)
	if err != nil {
		return err
	}

	_, err = c.db.CreateWeatherLog(ctx, database.CreateWeatherLogParams{
		Temperature: sql.NullFloat64{
			Float64: weatherRes.Current.Temperature2M,
			Valid:   true,
		},
		WindSpeed: sql.NullFloat64{
			Float64: weatherRes.Current.WindSpeed10M,
			Valid:   true,
		},
		CloudCover: sql.NullInt32{
			Int32: int32(weatherRes.Current.CloudCover),
			Valid: true,
		},
		Preassure: sql.NullFloat64{
			Float64: weatherRes.Current.PressureMsl,
			Valid:   true,
		},
		LocationID: location.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
