package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
)

func backfill(ctx context.Context, c *config, cmd command) error {
	backfillFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	days := backfillFlags.Int("days", 7, "specify how many days back the weahter logs should be backfilled")
	help := backfillFlags.Bool("help", false, "gives instructions about the backfill command")

	if err := backfillFlags.Parse(cmd.Args); err != nil {
		return err
	}

	if help != nil && *help {
		fmt.Println("usage: backfill [--days] [--help] <LOCATION>")
		return nil
	}

	args := backfillFlags.Args()

	if len(args) == 0 {
		return errors.New("not enough arguments provided")
	}

	locationName := args[0]

	if days != nil {
		if err := c.dbMocks.CreateWeatherLogsMocks(ctx, *days, locationName); err != nil {
			return fmt.Errorf("failed to generate the weather logs. %v", err)
		}

		log.Println("Successfully generated logs!")
	}

	return nil
}
