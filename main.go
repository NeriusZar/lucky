package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := &config{}
	if err := c.init(); err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-signalChan:
				log.Println("Got SIGINT/SIGTERM, exiting.")
				cancel()
				os.Exit(1)
			case <-ctx.Done():
				log.Println("Done.")
				os.Exit(1)
			}
		}
	}()

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	commands := supportedCommands()
	args := os.Args
	if len(args) < 2 {
		log.Fatal("no commands were provided")
	}
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	if err := commands.run(ctx, c, cmd); err != nil {
		log.Fatalf("failed to execute command. %v", err)
	}
}
