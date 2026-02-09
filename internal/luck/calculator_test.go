package luck

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSanpshotScore(t *testing.T) {
	randomUUID := uuid.New()
	now := time.Now()
	tests := map[string]struct {
		snap       FactorData
		locationId uuid.UUID
		factors    []Factor
		want       LuckResult
	}{
		"sweet spot": {
			snap: FactorData{
				Timestamp:     now,
				Temperature2M: ptr(20.0),
				WindSpeed10M:  ptr(5.0),
				CloudCover:    ptr(50),
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       100,
				Confidence:  100,
				FactorCount: 3,
				Timestamp:   now,
			},
		},
		"rough day": {
			snap: FactorData{
				Timestamp:     now,
				Temperature2M: ptr(10.0),
				WindSpeed10M:  ptr(50.0),
				CloudCover:    ptr(100),
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       18,
				Confidence:  100,
				FactorCount: 3,
				Timestamp:   now,
			},
		},
		"conflicting factors": {
			snap: FactorData{
				Timestamp:     now,
				Temperature2M: ptr(20.0),
				WindSpeed10M:  ptr(50.0),
				CloudCover:    ptr(30),
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       58,
				Confidence:  100,
				FactorCount: 3,
				Timestamp:   now,
			},
		},
		"missing weather data": {
			snap: FactorData{
				Timestamp:     now,
				Temperature2M: ptr(15.0),
				CloudCover:    ptr(0),
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       67,
				Confidence:  67,
				FactorCount: 2,
				Timestamp:   now,
			},
		},
		"wind only": {
			snap: FactorData{
				Timestamp:    now,
				WindSpeed10M: ptr(5.0),
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       100,
				Confidence:  33,
				FactorCount: 2,
				Timestamp:   now,
			},
		},
		"no data": {
			snap: FactorData{
				Timestamp: now,
			},
			locationId: randomUUID,
			factors:    []Factor{tempFactor{Name: "temperature factor"}, windFactor{Name: "wind factor"}, cloudCoverFactor{Name: "cloud cover factor"}},
			want: LuckResult{
				LocationId:  randomUUID,
				Score:       0,
				Confidence:  0,
				FactorCount: 3,
				Timestamp:   now,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lc := LuckCalculator{
				factors: test.factors,
			}
			result, err := lc.snapshotScore(test.snap, test.locationId)
			if err != nil {
				t.Error(err)
			}

			if result.Score != test.want.Score {
				t.Fatalf("expected score: %v, got %v", test.want.Score, result.Score)
			}

			if result.Confidence != test.want.Confidence {
				t.Fatalf("expected confidence: %v, got %v", test.want.Confidence, result.Confidence)
			}
		})
	}
}

type tempFactor struct {
	Name string
}

func (tf tempFactor) Score(data FactorData) (float64, error) {
	if data.Temperature2M == nil {
		return 0.0, ErrNoDataForFactor
	}

	if *data.Temperature2M >= 20 {
		return 1.0, nil
	} else if *data.Temperature2M >= 15 {
		return 0.7, nil
	} else {
		return 0.2, nil
	}
}

func (tf tempFactor) Weight() float64 {
	return 2
}

type windFactor struct {
	Name string
}

func (tf windFactor) Score(data FactorData) (float64, error) {
	if data.WindSpeed10M == nil {
		return 0.0, ErrNoDataForFactor
	}

	if *data.WindSpeed10M >= 50 {
		return 0.0, nil
	} else if *data.WindSpeed10M >= 20 {
		return 0.5, nil
	} else {
		return 1, nil
	}
}

func (tf windFactor) Weight() float64 {
	return 1.5
}

type cloudCoverFactor struct {
	Name string
}

func (tf cloudCoverFactor) Score(data FactorData) (float64, error) {
	if data.CloudCover == nil {
		return 0.0, ErrNoDataForFactor
	}

	if *data.CloudCover == 100 {
		return 0.4, nil
	} else if *data.CloudCover >= 50 {
		return 1, nil
	} else {
		return 0.6, nil
	}
}

func (tf cloudCoverFactor) Weight() float64 {
	return 1
}

func ptr[T any](v T) *T {
	return &v
}
