package cmd

import (
	"fmt"
)

func current(c *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("You have to provide location name")
	}

	locationName := args[0]
	location, ok := c.db.GetLocation(locationName)
	if !ok {
		return fmt.Errorf("There is no such location in database")
	}

	fmt.Println("Current weather information:")

	weatherInfo, err := c.api.GetCurrentAtmosphericData(location.Latitude, location.Longitude)
	if err != nil {
		return err
	}

	fmt.Printf("Temperature %.2f %s\n", weatherInfo.Current.Temperature2M, weatherInfo.CurrentUnits.Temperature2M)
	fmt.Printf("Wind %.2f %s\n", weatherInfo.Current.WindSpeed10M, weatherInfo.CurrentUnits.WindSpeed10M)
	fmt.Printf("Cloud cover %d %s\n", weatherInfo.Current.CloudCover, weatherInfo.CurrentUnits.CloudCover)
	fmt.Printf("Pressure %.2f %s\n", weatherInfo.Current.PressureMsl, weatherInfo.CurrentUnits.PressureMsl)

	return nil
}
