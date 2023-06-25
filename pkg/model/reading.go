package model

import "time"

// ElectricityReading defines the reading metadata
type ElectricityReading struct {
	Time    time.Time
	Reading float64
}
