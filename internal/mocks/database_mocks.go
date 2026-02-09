package mocks

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/NeriusZar/lucky/internal/database"
)

type DbMocks struct {
	db *database.Queries
}

func NewDbMocks(db *database.Queries) DbMocks {
	return DbMocks{
		db: db,
	}
}

func (m *DbMocks) CreateWeatherLogsMocks(ctx context.Context, days int, location string) error {
	m.cleanup(ctx)
	log.Printf("creating a %d days long data of hourly weather logs", days)

	loc, err := m.db.GetLocationByName(ctx, location)
	if err != nil {
		return fmt.Errorf("failed to get the location for weather logs. %v", err)
	}

	from := time.Now()
	for i := range 24 * days {
		updatedAt := from.Add(-(time.Duration(i) * time.Hour))

		m.db.CreateWeatherLogDebug(ctx, database.CreateWeatherLogDebugParams{
			CloudCover: sql.NullInt32{
				Int32: int32(rand.Float64() * 100.0),
				Valid: true,
			},
			Preassure: sql.NullFloat64{
				Float64: 980.0 + rand.Float64()*70,
				Valid:   true,
			},
			WindSpeed: sql.NullFloat64{
				Float64: rand.Float64() * 25,
				Valid:   true,
			},
			Temperature: sql.NullFloat64{
				Float64: rand.Float64() * 30,
				Valid:   true,
			},
			LocationID: loc.ID,
			UpdatedAt:  updatedAt,
		})
	}

	return nil
}

func (m *DbMocks) cleanup(ctx context.Context) {
	log.Println("cleaning up created logs...")
	m.db.RemoveDebugLogs(ctx)
}
