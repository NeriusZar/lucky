package cmd

import (
	"fmt"
	"os"
)

func exit(c *Config, args ...string) error {
	fmt.Println("Closing the Lucky... Goodbye!")
	os.Exit(0)
	return nil
}
