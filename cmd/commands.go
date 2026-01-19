package cmd

type command struct {
	name        string
	description string
	callback    func() error
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
	}
}
