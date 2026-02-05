package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/NeriusZar/lucky/internal/luck"
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

	location := cmd.Args[0]

	if daysCount == nil || *daysCount == 0 {
		return errors.New("invalid days parameter value. You can only provide values in range [1, 30]")
	}
	daysDuration := time.Duration(*daysCount) * time.Hour * 24
	to := time.Now()
	from := to.Add(-daysDuration)

	results, err := luck.DetermineLuck(ctx, location, from, to)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		return fmt.Errorf("Failed to gather any luck predictions for time period from %s to %s", from.Format("2006-02-10"), to.Format("2006-02-10"))
	}

	return nil
}
