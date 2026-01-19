package cmd

import "fmt"

func help() error {
	fmt.Println("Usage of Lucky:")
	for _, c := range getListOfCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
