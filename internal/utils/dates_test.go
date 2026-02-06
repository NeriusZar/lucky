package utils

import (
	"testing"
	"time"
)

func TestGetTimeRangeFromNow(t *testing.T) {
	loc := time.Now().Location()
	tests := map[string]struct {
		name      string
		daysCount int
		from      time.Time
		expected  time.Time
	}{
		"zero days case": {
			daysCount: 0,
			from:      time.Date(2026, 01, 1, 0, 0, 0, 0, loc),
			expected:  time.Date(2026, 01, 1, 0, 0, 0, 0, loc),
		},
		"yesterday": {
			daysCount: 1,
			from:      time.Date(2026, 01, 2, 0, 0, 0, 0, loc),
			expected:  time.Date(2026, 01, 1, 0, 0, 0, 0, loc),
		},
		"multiple days back": {
			daysCount: 10,
			from:      time.Date(2026, 01, 11, 0, 0, 0, 0, loc),
			expected:  time.Date(2026, 01, 1, 0, 0, 0, 0, loc),
		},
		"back across months/years": {
			daysCount: 1,
			from:      time.Date(2026, 01, 01, 0, 0, 0, 0, loc),
			expected:  time.Date(2025, 12, 31, 0, 0, 0, 0, loc),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := GetXDaysBack(tc.daysCount, tc.from)
			if !got.Equal(tc.expected) {
				t.Fatalf("expected: %#v, got: %#v", tc.expected, got)
			}
		})
	}
}
