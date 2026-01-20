package cmd

import (
	"fmt"
	"os"

	"github.com/NeriusZar/lucky/internal/meteoapi"
)

func exit(meteoapi meteoapi.ApiClient, args ...string) error {
	fmt.Println("Closing the Lucky... Goodbye!")
	os.Exit(0)
	return nil
}
