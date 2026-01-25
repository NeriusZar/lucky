package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NeriusZar/lucky/internal/database"
	"github.com/google/uuid"
)

func collectWeatherLogs(c *Config, args ...string) error {
	duration := time.Minute * 10
	ticker := time.NewTicker(duration)

	fmt.Println("Started collecting weather logs...")
	for ; ; <-ticker.C {
		fmt.Printf("Checking weather on %v\n", time.Now().UTC())
		locations, err := c.db.GetAllLocations(context.Background())
		if err != nil {
			fmt.Println("Failed to get locations")
			continue
		}

		for _, l := range locations {
			if err := fetchAndStoreWeatherLog(c, l); err != nil {
				fmt.Println("Failed to get locations")
			}
		}
	}

	return nil
}

func fetchAndStoreWeatherLog(c *Config, location database.Location) error {
	weatherRes, err := c.api.GetCurrentAtmosphericData(location.Latitude, location.Longitude)
	if err != nil {
		return err
	}

	_, err = c.db.CreateWeatherLog(context.Background(), database.CreateWeatherLogParams{
		ID: uuid.New(),
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
