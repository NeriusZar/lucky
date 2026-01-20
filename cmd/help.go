package cmd

import (
	"fmt"

	"github.com/NeriusZar/lucky/internal/meteoapi"
)

func help(meteoapi meteoapi.ApiClient, args ...string) error {
	fmt.Println("Usage of Lucky:")
	for _, c := range getListOfCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
