package events

import (
	"time"
)

type Event struct {
	AircraftID string
	Latitude   float64
	Longitude  float64
	StationID  string
	Timestamp  *time.Time
}
