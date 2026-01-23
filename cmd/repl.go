package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/NeriusZar/lucky/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func StartRepl() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getListOfCommands()

	dbUrl, ok := os.LookupEnv("POSTGRESQL_URL")
	if !ok {
		log.Fatal("db url was not provided in .env file")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("failed to open database")
	}
	dbQueries := database.New(db)
	config := NewConfig(dbQueries)

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
