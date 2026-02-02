package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/NeriusZar/lucky/internal/database"
	"github.com/NeriusZar/lucky/internal/meteoapi"
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

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					log.Println("Got SIGINT/SIGTERM, exiting.")
					cancel()
					os.Exit(1)
				case syscall.SIGHUP:
					log.Println("Got SIGHUP, reloading.")
					c.init()
				}
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

	if err := run(ctx, c, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	tick time.Duration
	api  *meteoapi.ApiClient
	db   *database.Queries
}

func (c *config) init() error {
	defaultTickParam, ok := os.LookupEnv("DEFAULT_TICK_SECONDS")
	if !ok {
		return fmt.Errorf("DEFAULT_TICK_SECONDS was not provided in .env file")
	}

	defaultTickInSeconds, err := strconv.Atoi(defaultTickParam)
	if err != nil {
		return fmt.Errorf("Failed to parse DEFAULT_TICK_SECONDS to integer")
	}

	dbUrl, ok := os.LookupEnv("POSTGRESQL_URL")
	if !ok {
		return fmt.Errorf("db url was not provided in .env file")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to open database")
	}
	dbQueries := database.New(db)

	api := meteoapi.NewApiClient()
	defaultTick := time.Duration(defaultTickInSeconds) * time.Second
	c.tick = defaultTick
	c.api = &api
	c.db = dbQueries

	return nil
}

func run(ctx context.Context, c *config, out io.Writer) error {
	if err := c.init(); err != nil {
		return fmt.Errorf("Failed to configure application. %v", err)
	}
	log.SetOutput(out)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.tick):
			fmt.Println("Doing job...")
			//TOOD do some operations
		}
	}
}
