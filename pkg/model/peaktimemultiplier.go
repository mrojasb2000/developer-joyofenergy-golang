package model

import "time"

// PeakTimeMultiplier defines the ... metadata
type PeakTimeMultiplier struct {
	DayOfWeek  time.Weekday
	Multiplier float64
}
