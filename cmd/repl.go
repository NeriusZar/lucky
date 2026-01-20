package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getListOfCommands()
	config := NewConfig()

	fmt.Println("Welcome to Lucky - your fishing luck detector.")
	for {
		fmt.Print("Lucky > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(&config, input[1:]...); err != nil {
			fmt.Printf("Failed to execute command %s", commandName)
			continue
		}
	}
}
