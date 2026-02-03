package main

import (
	"context"
	"errors"
)

type commands struct {
	registered map[string]func(context.Context, *config, command) error
}

type command struct {
	Name string
	Args []string
}

func supportedCommands() commands {
	return commands{
		registered: map[string]func(context.Context, *config, command) error{
			"collect": collectWeahterLogs,
			"add": addLocation,
			"locations": locations,
		},
	}
}

func (cmds commands) run(ctx context.Context, c *config, cmd command) error {
	callback, ok := cmds.registered[cmd.Name]
	if !ok {
		return errors.New("command does not exist")
	}

	if err := callback(ctx, c, cmd); err != nil {
		return err
	}

	return nil
}
