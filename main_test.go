package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRunConfigInitialization(t *testing.T) {
	setupAndTimeoutWrap(time.Second*1, func(ctx context.Context) {
		c := config{}

		err := run(ctx, &c, os.Stdout)
		if err != nil {
			t.Error(err)
		}

		if c.tick != 5*time.Second {
			t.Fail()
		}
	})
}

func setupAndTimeoutWrap(timeout time.Duration, test func(context.Context)) {
	func() {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, timeout)
		_ = godotenv.Load()
		defer cancel()

		test(ctx)
	}()
}
