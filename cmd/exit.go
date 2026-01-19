package cmd

import (
	"fmt"
	"os"
)

func exit() error {
	fmt.Println("Closing the Lucky... Goodbye!")
	os.Exit(0)
	return nil
}
