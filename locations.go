package main

import (
	"context"
	"flag"
	"fmt"
)

func locations(ctx context.Context, c *config, cmd command) error {
	addFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	help := addFlags.Bool("help", false, "gives instructions about the add command")
	withCoords := addFlags.Bool("c", false, "displays locations with coordinates")

	if err := addFlags.Parse(cmd.Args); err != nil {
		return fmt.Errorf("failed to parse flags %v", err)
	}

	if help != nil && *help {
		fmt.Println("usage: locations [-c]")
		return nil
	}

	locations, err := c.db.GetAllLocations(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve locations from database")
	}

	for _, l := range locations {
		if withCoords != nil && *withCoords {
			fmt.Printf("- %s latitude: %v longitude %v\n", l.Name, l.Latitude, l.Longitude)
		} else {
			fmt.Printf("- %s\n", l.Name)
		}
	}

	return nil
}
