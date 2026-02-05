package luck

import (
	"context"
	"time"
)

type Luck struct {
	location    string
	score       float64
	confidence  float64
	factorCount int
}

func DetermineLuck(ctx context.Context, location string, from, to time.Time) ([]Luck, error) {
	return []Luck{}, nil
}
