package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var lat, long float64
	var err error
	for {
		fmt.Println("Enter the latitude of the location:")
		fmt.Printf("Lucky > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Failed to capture the latitude")
			continue
		}

		lat, err = strconv.ParseFloat(input[0], 64)
		if err != nil {
			fmt.Println("Failed to parse the provided value to float")
			continue
		}

		break
	}

	for {
		fmt.Println("Enter the longitude of the location:")
		fmt.Printf("Lucky > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Failed to capture the latitude")
			continue
		}

		long, err = strconv.ParseFloat(input[0], 64)
		if err != nil {
			fmt.Println("Failed to parse the provided value to float")
			continue
		}

		break
	}

	c.db.AddNewLocation(lat, long, name)

	fmt.Println("Successfully added!")

	return nil
}
