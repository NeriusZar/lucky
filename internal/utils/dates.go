package utils

import "time"

func GetXDaysBack(daysBack int, from time.Time) time.Time {
	d := time.Duration(daysBack) * time.Hour * 24
	return from.Add(-d)
}
