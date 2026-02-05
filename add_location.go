package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/NeriusZar/lucky/internal/database"
)

func addLocation(ctx context.Context, c *config, cmd command) error {
	addFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	help := addFlags.Bool("help", false, "gives instructions about the add command")

	if err := addFlags.Parse(cmd.Args); err != nil {
		return fmt.Errorf("failed to parse flags %v", err)
	}

	if help != nil && *help {
		fmt.Println("usage: add <name> <latitude> <longitude>")
		return nil
	}

	if len(cmd.Args) != 3 {
		return errors.New("not enough arguments provided")
	}

	name := cmd.Args[0]
	lat, err := strconv.ParseFloat(cmd.Args[1], 64)
	if err != nil {
		return fmt.Errorf("failed to parse latitude")
	}
	lon, err := strconv.ParseFloat(cmd.Args[2], 64)
	if err != nil {
		return fmt.Errorf("failed to parse longitude")
	}

	location, err := c.db.CreateLocation(ctx, database.CreateLocationParams{
		Name:      name,
		Latitude:  lat,
		Longitude: lon,
	})
	if err != nil {
		return fmt.Errorf("failed to add the location to the database. %v", err)
	}

	log.Printf("successfully added new location - %s", location.Name)

	return nil
}
