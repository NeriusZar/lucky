package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/NeriusZar/lucky/internal/utils"
)

func historicalScore(ctx context.Context, c *config, cmd command) error {
	scoreFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	help := scoreFlags.Bool("help", false, "gives instructions on how to use the score command")
	daysCount := scoreFlags.Int("d", 1, "the amout of days back from today to show historic data")

	if err := scoreFlags.Parse(cmd.Args); err != nil {
		return fmt.Errorf("failed to parse flags %v", err)
	}

	if help != nil && *help {
		fmt.Println("usage: score <location> [-d]")
		return nil
	}

	if len(cmd.Args) < 1 {
		return errors.New("not enough arguments provided")
	}

	locationName := cmd.Args[0]

	if daysCount == nil || *daysCount == 0 {
		return errors.New("invalid days parameter value. You can only provide values in range [1, 30]")
	}
	to := time.Now()
	from := utils.GetXDaysBack(*daysCount, to)

	location, err := c.db.GetLocationByName(ctx, locationName)
	if err != nil {
		return fmt.Errorf("Failed to get location with specified name. %v", err)
	}

	results, err := c.lc.DetermineLuck(ctx, location.ID, from, to)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return fmt.Errorf("Failed to gather any luck predictions for time period from %v to %v", from.Format(time.RFC3339), to.Format(time.RFC3339))
	}

	fmt.Printf("Fishing luck analysis for %s:\n", locationName)
	for _, luck := range results {
		fmt.Printf("Time: %s \n", luck.Timestamp.Format(time.DateTime))
		fmt.Printf("Luck score: %v\n", luck.Score)
		fmt.Printf("Confidence: %v\n", luck.Confidence)
		fmt.Printf("Factors included: %v\n", luck.FactorCount)
		fmt.Println("----------------------------------")
	}

	return nil
}
