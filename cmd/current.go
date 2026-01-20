package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/NeriusZar/lucky/internal/meteoapi"
)

func current(api meteoapi.ApiClient, args ...string) error {
	scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Println("Current weather information:")

	weatherInfo, err := api.GetCurrentAtmosphericData(lat, long)
	fmt.Printf("Temperature %.2f %s\n", weatherInfo.Current.Temperature2M, weatherInfo.CurrentUnits.Temperature2M)
	fmt.Printf("Wind %.2f %s\n", weatherInfo.Current.WindSpeed10M, weatherInfo.CurrentUnits.WindSpeed10M)
	fmt.Printf("Cloud cover %d %s\n", weatherInfo.Current.CloudCover, weatherInfo.CurrentUnits.CloudCover)
	fmt.Printf("Pressure %.2f %s\n", weatherInfo.Current.PressureMsl, weatherInfo.CurrentUnits.PressureMsl)

	return nil
}
