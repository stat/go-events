package models

import (
	"errors"
	"time"
)

// for additional ADS-B tranmission fields:
// https://www.flightradar24.com/blog/ads-b/#:~:text=ADS%2DB%20In%20Reception%3A%20ADS,%2C%20speed%2C%20and%20other%20information.

type ADSB struct {
	AircraftID string     `json:"aircraftID" validate:"required"`
	Latitude   float64    `json:"latitude" validate:"latitude"`
	Longitude  float64    `json:"longitude" validate:"longitude"`
	StationID  string     `json:"stationID" validate:"required"`
	Timestamp  *time.Time `json:"timestamp" validate:"required"`
}

var (
	ADSBValidateAircraftIDError = errors.New("ADSB event aircarft id validation error")
	ADSBValidateLatitudeError   = errors.New("ADSB event latitude validation error")
	ADSBValidateLongitudeError  = errors.New("ADSB event longitude validation error")
	ADSBValidateStationIDError  = errors.New("ADSB event station id validation error")
	ADSBValidateTimestampError  = errors.New("ADSB event timestamp validation error")
)

func (event *ADSB) Validate() error {
	if event.AircraftID == "" {
		return ADSBValidateAircraftIDError
	}

	if event.Latitude == 0 {
		return ADSBValidateLatitudeError
	}

	if event.Longitude == 0 {
		return ADSBValidateLongitudeError
	}

	if event.StationID == "" {
		return ADSBValidateStationIDError
	}

	if event.Timestamp == nil {
		return ADSBValidateTimestampError
	}

	return nil
}
