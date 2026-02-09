package luck

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"

	"github.com/NeriusZar/lucky/internal/database"
)

type LuckCalculator struct {
	db      *database.Queries
	factors []Factor
}

func NewLuckCalculator(db *database.Queries) LuckCalculator {
	return LuckCalculator{
		db: db,
	}
}

type LuckResult struct {
	LocationId  uuid.UUID
	Score       float64
	Confidence  float64
	Timestamp   time.Time
	FactorCount int
}

type FactorDataInterval int

const (
	daily  FactorDataInterval = 1
	hourly FactorDataInterval = 2
)

func (lc LuckCalculator) DetermineLuck(ctx context.Context, locationId uuid.UUID, from, to time.Time, isDaily bool) ([]LuckResult, error) {
	interval := hourly
	if isDaily {
		interval = daily
	}
	factorData, err := lc.getDataForPeriod(ctx, locationId, from, to, interval)
	if err != nil {
		return nil, fmt.Errorf("failed to gather luck factor data. %v", err)
	}

	results := make([]LuckResult, 0, len(factorData))
	for _, snap := range factorData {
		r, err := lc.snapshotScore(snap, locationId)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate snapshot score of time - %v. %v", snap.Timestamp, err)
		}
		results = append(results, r)
	}

	return results, nil
}

func (lc LuckCalculator) snapshotScore(snap FactorData, locationId uuid.UUID) (LuckResult, error) {
	var weightedScore, totalWeight, availableWeight float64
	factorsIncluded := 0
	for _, f := range lc.factors {
		totalWeight += f.Weight()
		score, err := f.Score(snap)
		if err != nil {
			if errors.Is(err, ErrNoDataForFactor) {
				continue
			} else {
				return LuckResult{}, nil
			}
		}
		weightedScore += (score * f.Weight())
		availableWeight += f.Weight()
	}

	var finalScore, confidence float64
	if availableWeight != 0 && totalWeight != 0 {
		finalScore = math.Round((weightedScore / availableWeight) * 100)
		confidence = math.Round((availableWeight / totalWeight) * 100)
	}

	return LuckResult{
		Score:       finalScore,
		LocationId:  locationId,
		Timestamp:   snap.Timestamp,
		FactorCount: factorsIncluded,
		Confidence:  confidence,
	}, nil
}

func (lc LuckCalculator) getDataForPeriod(ctx context.Context, locationId uuid.UUID, from, to time.Time, interval FactorDataInterval) ([]FactorData, error) {
	switch interval {
	case daily:
		dailyLogs, err := lc.db.GetDailyLogsWithinRange(ctx, database.GetDailyLogsWithinRangeParams{
			ID:         locationId,
			RangeStart: from,
			RangeEnd:   to,
		})
		if err != nil {
			return nil, err
		}

		results := make([]FactorData, 0, len(dailyLogs))
		for _, dl := range dailyLogs {
			cloudCover := int(dl.CloudCover)
			results = append(results, FactorData{
				Timestamp:     dl.DailyBucket,
				WindSpeed10M:  &dl.Speed,
				Temperature2M: &dl.Temperature,
				PressureMsl:   &dl.Preassure,
				CloudCover:    &cloudCover,
			})
		}
		return results, nil
	case hourly:
		hourlyLogs, err := lc.db.GetHourlyLogsWithinRange(ctx, database.GetHourlyLogsWithinRangeParams{
			ID:         locationId,
			RangeStart: from,
			RangeEnd:   to,
		})
		if err != nil {
			return nil, err
		}

		results := make([]FactorData, 0, len(hourlyLogs))
		for _, hl := range hourlyLogs {
			cloudCover := int(hl.CloudCover)
			results = append(results, FactorData{
				Timestamp:     hl.HourlyBucket,
				WindSpeed10M:  &hl.Speed,
				Temperature2M: &hl.Temperature,
				PressureMsl:   &hl.Preassure,
				CloudCover:    &cloudCover,
			})
		}

		populateHistoricalPreassure(results)
		return results, nil
	default:
		return nil, errors.New("unsupported time interval type")
	}
}

func populateHistoricalPreassure(snapshots []FactorData) {
	snapMap := make(map[time.Time]FactorData)
	for _, s := range snapshots {
		snapMap[s.Timestamp] = s
	}

	for i := range snapshots {
		targetTime := snapshots[i].Timestamp.Add(-(time.Hour * 3))
		if tSnap, ok := snapMap[targetTime]; ok {
			snapshots[i].PreassureMsl3hBefore = tSnap.PressureMsl
		}
	}
}
