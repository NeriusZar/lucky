package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/NeriusZar/lucky/internal/database"
	"github.com/google/uuid"
)

func addLocation(c *Config, args ...string) error {
	scanner := bufio.NewScanner(os.Stdin)

	var name string
	for {
		fmt.Println("Enter the name of the location:")
		fmt.Printf("Lucky > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Failed to capture the name")
			continue
		}

		name = input[0]
		break
	}

	fmt.Println("Enter the latitude of the location:")
	lat, err := getSingleFloatParameter()
	if err != nil {
		return err
	}

	fmt.Println("Enter the longitude of the location:")
	long, err := getSingleFloatParameter()
	if err != nil {
		return err
	}

	location, err := c.db.CreateLocation(context.Background(), database.CreateLocationParams{
		ID:        uuid.New(),
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	})
	if err != nil {
		return fmt.Errorf("Failed to add the location to the database. %v", err)
	}

	fmt.Printf("Successfully added %s!\n", location.Name)

	return nil
}

func getSingleFloatParameter() (float64, error) {
	scanner := bufio.NewScanner(os.Stdin)

	var param float64
	var err error
	for {
		fmt.Printf("Lucky > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Failed to capture the parameter")
			continue
		}

		param, err = strconv.ParseFloat(input[0], 64)
		if err != nil {
			fmt.Println("Failed to parse the provided value to float")
			continue
		}

		break
	}

	return param, nil
}
