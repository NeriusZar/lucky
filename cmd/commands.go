package cmd

import "github.com/NeriusZar/lucky/internal/meteoapi"

type command struct {
	name        string
	description string
	callback    func(meteoapi.ApiClient, ...string) error
}

func getListOfCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays detailed information on all the supported commands",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "Exits the app safely",
			callback:    exit,
		},
		"current": {
			name:        "current",
			description: "Retrieves latest information about current weather conditions in a specific area",
			callback:    current,
		},
	}
}
