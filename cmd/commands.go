package cmd

type command struct {
	name        string
	description string
	callback    func(*Config, ...string) error
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
		"add_location": {
			name:        "add_location",
			description: "Adds a new location to the database",
			callback:    addLocation,
		},
		"locations": {
			name:        "locations",
			description: "Displays a list of all added locations",
			callback:    locations,
		},
	}
}
