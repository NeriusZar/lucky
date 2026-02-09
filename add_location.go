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
		fmt.Println("usage: add <NAME> <LATITUDE> <LONGITUDE>")
		return nil
	}

	args := addFlags.Args()
	if len(args) != 3 {
		return errors.New("not enough arguments provided")
	}

	name := args[0]
	lat, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("failed to parse latitude")
	}
	lon, err := strconv.ParseFloat(args[2], 64)
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
