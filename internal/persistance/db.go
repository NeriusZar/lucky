package persistance

import (
	"sync"

	"github.com/NeriusZar/lucky/internal/models"
)

type LuckyDB struct {
	mu        sync.Mutex
	locations map[string]models.Location
}

func NewLuckyDB() LuckyDB {
	return LuckyDB{
		mu:        sync.Mutex{},
		locations: map[string]models.Location{},
	}
}

func (db *LuckyDB) AddNewLocation(lat, long float64, name string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.locations[name] = models.Location{
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	}
}

func (db *LuckyDB) GetLocation(name string) (models.Location, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	location, ok := db.locations[name]

	return location, ok
}

func (db *LuckyDB) GetAllLocations() []models.Location {
	locations := []models.Location{}
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, v := range db.locations {
		locations = append(locations, v)
	}
	return locations
}
