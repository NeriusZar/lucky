package luck

import (
	"errors"
	"time"
)

type Factor interface {
	Score(data FactorData) (float64, error)
	Weight() float64
}

type FactorData struct {
	Timestamp     time.Time
	Temperature2M float64
	WindSpeed10M  float64
	CloudCover    int
	PressureMsl   float64
}

var ErrNoDataForFactor = errors.New("No data provided for the score calculations")
